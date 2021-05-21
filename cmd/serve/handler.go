package serve

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	gen "gitlab.warungpintar.co/farhan.ramadhan/onboard-service/proto"
	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/mux"
	"gitlab.warungpintar.co/farhan.ramadhan/onboard-service/internal/services"
)

type routeHandler struct {
	messageSvc        services.MessageServiceInterface
	messageGrpcClient gen.MessageServiceClient
}

func NewRouteHandler(messageSvc services.MessageServiceInterface) *routeHandler {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	client := gen.NewMessageServiceClient(conn)

	log.Println("GRPC Connected To :8081")

	return &routeHandler{
		messageSvc:        messageSvc,
		messageGrpcClient: client,
	}
}

func (rh *routeHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["key"]

	msg := "World"
	if len(keys) > 0 {
		msg = (keys[0])
	}

	payload := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{

		Status:  "Success",
		Message: "Hello " + msg,
	}

	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (rh *routeHandler) InsertMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	message := params["message"]

	_, err := rh.messageGrpcClient.InsertMessage(context.Background(), &gen.MessageRequest{
		Body: message,
	})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
}

func (rh *routeHandler) GetAllMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := rh.messageGrpcClient.GetAllMessages(context.Background(), &empty.Empty{})
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var data struct {
		Data struct {
			Length   int `json:"length"`
			Messages []struct {
				Body string `json:"body"`
			} `json:"messages"`
		} `json:"data"`
	}

	for _, v := range messages.Body {
		var message struct {
			Body string `json:"body"`
		}

		message.Body = v.Body

		data.Data.Messages = append(data.Data.Messages, message)
	}

	data.Data.Length = len(data.Data.Messages)

	message, err := json.Marshal(data)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(message)
}
