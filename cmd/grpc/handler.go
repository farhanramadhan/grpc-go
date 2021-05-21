package grpc

import (
	"context"

	gen "gitlab.warungpintar.co/farhan.ramadhan/onboard-service/proto"

	"github.com/golang/protobuf/ptypes/empty"
	"gitlab.warungpintar.co/farhan.ramadhan/onboard-service/internal/services"
)

type grpcHandler struct {
	messageSvc services.MessageServiceInterface
}

func NewGrpcHandler(messageService services.MessageServiceInterface) *grpcHandler {
	return &grpcHandler{
		messageSvc: messageService,
	}
}

func (gh *grpcHandler) InsertMessage(ctx context.Context, in *gen.MessageRequest) (*gen.MessageResponse, error) {
	err := gh.messageSvc.InsertMessage(ctx, in.Body)
	if err != nil {
		return nil, err
	}

	res := &gen.MessageResponse{
		Body: in.Body,
	}

	return res, nil
}

func (gh *grpcHandler) GetAllMessages(ctx context.Context, in *empty.Empty) (*gen.MessagesResponse, error) {
	messages, err := gh.messageSvc.GetAllMessages(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*gen.MessageResponse, 0)
	for _, message := range messages {
		messageRes := &gen.MessageResponse{
			Body: message.Body,
		}
		res = append(res, messageRes)
	}

	messagesResponse := &gen.MessagesResponse{
		Body: res,
	}

	return messagesResponse, nil
}
