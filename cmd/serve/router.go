package serve

import "github.com/gorilla/mux"

type Router struct {
	Router *mux.Router
}

func NewRouter(handler *routeHandler) *Router {
	r := mux.NewRouter()

	r.HandleFunc("/healthz", handler.HealthCheck).Methods("GET")
	r.HandleFunc("/message/{message}", handler.InsertMessage).Methods("GET")
	r.HandleFunc("/message", handler.GetAllMessages).Methods("GET")

	return &Router{
		Router: r,
	}
}
