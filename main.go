package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"gitlab.warungpintar.co/farhan.ramadhan/onboard-service/proto"
	"google.golang.org/grpc/reflection"

	"gitlab.warungpintar.co/farhan.ramadhan/onboard-service/cmd/grpc"

	"gitlab.warungpintar.co/farhan.ramadhan/onboard-service/internal/messages/repository"
	"gitlab.warungpintar.co/farhan.ramadhan/onboard-service/internal/services"

	"gitlab.warungpintar.co/farhan.ramadhan/onboard-service/cmd/serve"
)

func main() {
	startServer()
}

func startServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	repo := repository.NewMessageRepository()

	svc := services.NewMessageService(repo)

	handler := serve.NewRouteHandler(svc)

	router := serve.NewRouter(handler)

	go startGRPCServer(svc)

	log.Println("Starting server on port ", port)

	log.Fatal(http.ListenAndServe(":"+port, router.Router))
}

func startGRPCServer(svc services.MessageServiceInterface) {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "8081"
	}

	handler := grpc.NewGrpcHandler(svc)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewGrpcServer(*handler)

	reflection.Register(grpcServer.Server)

	proto.RegisterMessageServiceServer(grpcServer.Server, handler)

	log.Println("Starting GRPC Server on port ", port)

	grpcServer.Server.Serve(lis)
}
