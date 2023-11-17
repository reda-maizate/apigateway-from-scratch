package grpc

import (
	"api-gateway/internal/core/domain"
	"api-gateway/internal/core/services"
)

type GrpcNoteHandler struct {
	svc services.NoteService
}

func NewGrpcNoteHandler(svc services.NoteService) *GrpcNoteHandler {
	return &GrpcNoteHandler{
		svc: svc,
	}
}

func (gnh *GrpcNoteHandler) Create(title, content string) error {
	return gnh.svc.Create(title, content)
}

func (gnh *GrpcNoteHandler) GetAll() ([]*domain.Note, error) {
	return gnh.svc.GetAll()
}
