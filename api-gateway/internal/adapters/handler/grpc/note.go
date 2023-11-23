package grpc

import (
	"api-gateway/internal/core/ports"
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

func (h *GrpcNoteHandler) Create(createNoteParams ports.CreateNoteParams) error {
	return h.noteService.Create(createNoteParams)
}

func (h *GrpcNoteHandler) GetAll() (ports.GetAllResponse, error) {
	return h.noteService.GetAll()
}
