package repository

import (
	"context"

	"gitlab.warungpintar.co/farhan.ramadhan/onboard-service/internal/messages"
)

func (mr *messageRepository) InsertMessage(ctx context.Context, message *messages.Message) (err error) {
	mr.db = append(mr.db, message)
	return nil
}

func (mr *messageRepository) GetAllMessages(ctx context.Context) (messages []*messages.Message, err error) {
	return mr.db, nil
}
