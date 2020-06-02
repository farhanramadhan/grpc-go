package services

import (
	"context"
	"time"

	"gitlab.warungpintar.co/farhan.ramadhan/onboard-service/internal/messages/repository"

	"gitlab.warungpintar.co/farhan.ramadhan/onboard-service/internal/messages"
)

type MessageServiceInterface interface {
	InsertMessage(ctx context.Context, body string) (err error)
	GetAllMessages(ctx context.Context) (messages []*messages.Message, err error)
}

type messageService struct {
	messageRepo repository.MessageRepositoryInterface
}

func NewMessageService(messageRepo repository.MessageRepositoryInterface) *messageService {
	return &messageService{
		messageRepo: messageRepo,
	}
}

func (ms *messageService) InsertMessage(ctx context.Context, body string) (err error) {
	message := &messages.Message{
		Body:      body,
		CreatedAt: time.Now(),
	}

	err = ms.messageRepo.InsertMessage(ctx, message)
	if err != nil {
		return err
	}

	return
}

func (ms *messageService) GetAllMessages(ctx context.Context) (messages []*messages.Message, err error) {
	messages, err = ms.messageRepo.GetAllMessages(ctx)
	if err != nil {
		return nil, err
	}

	return messages, err
}
