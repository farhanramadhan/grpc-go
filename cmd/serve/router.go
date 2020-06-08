package serve

import "github.com/gorilla/mux"

type Router struct {
	Router *mux.Router
}

func NewRouter(handler *routeHandler) *Router {
	r := mux.NewRouter()

	routePrefix := r.PathPrefix("/farhan-onboard-http").Subrouter()
	routePrefix.HandleFunc("/healthz", handler.HealthCheck).Methods("GET")
	routePrefix.HandleFunc("/message/{message}", handler.InsertMessage).Methods("GET")
	routePrefix.HandleFunc("/message", handler.GetAllMessages).Methods("GET")

	return &Router{
		Router: r,
	}
}
