package app

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/a13hander/auth-service-api/internal/app/auth_v1/interceptors"
	"github.com/a13hander/auth-service-api/internal/config"
	"github.com/a13hander/auth-service-api/internal/service/metric"
	descAccess "github.com/a13hander/auth-service-api/pkg/access_v1"
	descAuth "github.com/a13hander/auth-service-api/pkg/auth_v1"
	_ "github.com/a13hander/auth-service-api/statik"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider  *serviceProvider
	config           *config.Config
	grpcV1Server     *grpc.Server
	httpServer       *http.Server
	swaggerServer    *http.Server
	prometheusServer *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.closeAll()
		closer.wait()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(4)

	go func() {
		defer wg.Done()

		err := a.runGrpcV1Server(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runHttpServer(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runSwaggerServer(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runPrometheusServer()
		if err != nil {
			log.Fatalf("failed to run Prometheus server: %v", err)
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		metric.Init,
		a.initServiceProvider,
		a.initGrpcV1Server,
		a.initHttpServer,
		a.initSwaggerServer,
		a.initPrometheusServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	a.config = config.GetConfig()
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGrpcV1Server(ctx context.Context) error {
	transportCreds, err := credentials.NewServerTLSFromFile("service.pem", "service.key")
	if err != nil {
		return err
	}

	rateLimiterInterceptor := interceptors.NewRateLimiterInterceptor(a.serviceProvider.GetRateLimiter(ctx))

	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "auth-service",
		MaxRequests: 3,
		Timeout:     5 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return failureRatio >= 0.6
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf("Circuit Breaker: %s, changed from %v, to %v\n", name, from, to)
		},
	})
	circuitBreakerInterceptor := interceptors.NewCircuitBreakerInterceptor(cb)

	a.grpcV1Server = grpc.NewServer(
		grpc.Creds(transportCreds),
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				interceptors.ErrorCodesInterceptor,
				interceptors.ValidateInterceptor,
				rateLimiterInterceptor.Unary,
				interceptors.MetricsInterceptor,
				circuitBreakerInterceptor.Unary,
			),
		),
	)
	reflection.Register(a.grpcV1Server)

	descAuth.RegisterAuthV1Server(a.grpcV1Server, a.serviceProvider.GetAuthV1ServerImpl(ctx))
	descAccess.RegisterAccessV1Server(a.grpcV1Server, a.serviceProvider.GetAccessV1ServerImpl(ctx))

	return nil
}

func (a *App) runGrpcV1Server(_ context.Context) error {
	log.Printf("Grpc server starting on %s\n", a.config.GrpcPort)

	listener, err := net.Listen("tcp", a.config.GrpcPort)
	if err != nil {
		return err
	}

	err = a.grpcV1Server.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initHttpServer(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := descAuth.RegisterAuthV1HandlerFromEndpoint(ctx, mux, a.config.GrpcPort, opts)
	if err != nil {
		return err
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	a.httpServer = &http.Server{
		Addr:    a.config.HttpPort,
		Handler: corsMiddleware.Handler(mux),
	}

	return nil
}

func (a *App) initSwaggerServer(_ context.Context) error {
	fileSystem, err := fs.New()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(fileSystem)))
	mux.HandleFunc("/api.swagger.json", serveSwaggerFile("/api.swagger.json"))

	a.swaggerServer = &http.Server{
		Addr:    a.config.SwaggerPort,
		Handler: mux,
	}

	return nil
}

func (a *App) runHttpServer(_ context.Context) error {
	log.Printf("Http server starting on %s\n", a.httpServer.Addr)

	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runSwaggerServer(_ context.Context) error {
	log.Printf("Swagger server starting on %s\n", a.swaggerServer.Addr)

	err := a.swaggerServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func serveSwaggerFile(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fileSystem, err := fs.New()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		file, err := fileSystem.Open(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer func() {
			_ = file.Close()
		}()

		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (a *App) initPrometheusServer(_ context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	a.prometheusServer = &http.Server{
		Addr:    "localhost:2112",
		Handler: mux,
	}

	return nil
}

func (a *App) runPrometheusServer() error {
	err := a.prometheusServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
