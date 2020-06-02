package main

import (
	"log"
	"net/http"
	"os"

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

	log.Println("Starting server on port ", port)

	log.Fatal(http.ListenAndServe(":"+port, router.Router))
}
