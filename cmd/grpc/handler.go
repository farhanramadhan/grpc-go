package grpc

import (
	"context"

	"gitlab.warungpintar.co/farhan.ramadhan/onboard-service/proto"

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

func (gh *grpcHandler) InsertMessage(ctx context.Context, in *proto.MessageRequest) (*proto.MessageResponse, error) {
	err := gh.messageSvc.InsertMessage(ctx, in.Body)
	if err != nil {
		return nil, err
	}

	res := &proto.MessageResponse{
		Body: in.Body,
	}

	return res, nil
}

func (gh *grpcHandler) GetAllMessages(ctx context.Context, in *empty.Empty) (*proto.MessagesResponse, error) {
	messages, err := gh.messageSvc.GetAllMessages(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]*proto.MessageResponse, 0)
	for _, message := range messages {
		messageRes := &proto.MessageResponse{
			Body: message.Body,
		}
		res = append(res, messageRes)
	}

	messagesResponse := &proto.MessagesResponse{
		Body: res,
	}

	return messagesResponse, nil
}
