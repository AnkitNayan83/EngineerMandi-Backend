package services

import "github.com/google/uuid"

type SocketService interface {
	SaveMessage(conversationId uuid.UUID, senderId uuid.UUID, content string, attachment string, attachmentType string) error
}

type socketService struct {
}
