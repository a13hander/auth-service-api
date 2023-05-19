package app

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/a13hander/auth-service-api/internal/app/grpc_v1/interceptors"
	"github.com/a13hander/auth-service-api/internal/config"
	desc "github.com/a13hander/auth-service-api/pkg/auth_v1"
	_ "github.com/a13hander/auth-service-api/statik"
)

type App struct {
	serviceProvider *serviceProvider
	config          *config.Config
	grpcV1Server    *grpc.Server
	httpV1Server    *http.Server
	swaggerServer   *http.Server
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
	wg.Add(3)

	go func() {
		defer wg.Done()

		err := a.runGrpcV1Server(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runHttpV1Server(ctx)
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

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGrpcV1Server,
		a.initHttpV1Server,
		a.initSwaggerServer,
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
	a.grpcV1Server = grpc.NewServer(grpc.UnaryInterceptor(interceptors.ValidateInterceptor))
	reflection.Register(a.grpcV1Server)
	desc.RegisterAuthV1Server(a.grpcV1Server, a.serviceProvider.GetGrpcV1ServerImpl(ctx))

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

func (a *App) initHttpV1Server(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := desc.RegisterAuthV1HandlerFromEndpoint(ctx, mux, a.config.GrpcPort, opts)
	if err != nil {
		return err
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	a.httpV1Server = &http.Server{
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

func (a *App) runHttpV1Server(_ context.Context) error {
	log.Printf("Http server starting on %s\n", a.httpV1Server.Addr)

	err := a.httpV1Server.ListenAndServe()
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
		defer file.Close()

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
