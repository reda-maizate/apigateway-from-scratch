package grpc

import (
	"api-gateway/internal/core/domain"
	"api-gateway/internal/core/services"
)

type GrpcNoteHandler struct {
	noteService services.NoteService
}

func NewGrpcNoteHandler(noteService services.NoteService) *GrpcNoteHandler {
	return &GrpcNoteHandler{
		noteService: noteService,
	}
}

func (h *GrpcNoteHandler) Create(title, content, userUuid string) error {
	return h.noteService.Create(title, content, userUuid)
}

func (h *GrpcNoteHandler) GetAll() ([]*domain.Note, error) {
	return h.noteService.GetAll()
}
