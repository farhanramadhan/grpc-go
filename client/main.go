package main

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"gitlab.warungpintar.co/farhan.ramadhan/onboard-service/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := proto.NewMessageServiceClient(conn)

	message, err := client.InsertMessage(context.Background(), &proto.MessageRequest{
		Body: "test1",
	})
	if err != nil {
		log.Println(err)
	}
	log.Println("Messages 1 :", message)

	message, err = client.InsertMessage(context.Background(), &proto.MessageRequest{
		Body: "test2",
	})
	if err != nil {
		log.Println(err)
	}
	log.Println("Messages 2 :", message)

	messages, err := client.GetAllMessages(context.Background(), &empty.Empty{})
	if err != nil {
		log.Println(err)
	}
	log.Println("Messagess :", messages)
}
