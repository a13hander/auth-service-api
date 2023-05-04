package main

import (
	"log"
	"net"

	"github.com/a13hander/auth-service-api/internal/container"
	desc "github.com/a13hander/auth-service-api/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cont := container.Build()
	config := container.GetConfig()

	listener, err := net.Listen("tcp", config.GrpcPort)
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	desc.RegisterAuthV1Server(server, cont.Grpc.V1)

	err = server.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}
