package services

import (
	"context"
	"testing"

	"gitlab.warungpintar.co/farhan.ramadhan/onboard-service/internal/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestInsertMessage(t *testing.T) {
	// Init Dependency
	ctrl := gomock.NewController(t)
	repo := mock.NewMockMessageRepositoryInterface(ctrl)
	svc := NewMessageService(repo)

	// Expected Call
	repo.EXPECT().InsertMessage(gomock.Any(), gomock.Any()).
		Times(1).Return(nil)

	err := svc.InsertMessage(context.Background(), "testing")

	assert.NoError(t, err)
}
