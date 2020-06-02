package repository

import (
	"context"

	"gitlab.warungpintar.co/farhan.ramadhan/onboard-service/internal/messages"
)

type MessageRepositoryInterface interface {
	InsertMessage(ctx context.Context, message *messages.Message) (err error)
	GetAllMessages(ctx context.Context) (messages []*messages.Message, err error)
}

type messageRepository struct {
	db []*messages.Message
}

func NewMessageRepository() *messageRepository {
	return &messageRepository{
		db: make([]*messages.Message, 0),
	}
}
