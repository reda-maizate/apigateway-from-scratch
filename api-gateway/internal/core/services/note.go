package services

import (
	"api-gateway/internal/core/ports"
)

type NoteService struct {
	noteRepository ports.NoteRepository
}

func NewNoteService(repo ports.NoteRepository) *NoteService {
	return &NoteService{
		noteRepository: repo,
	}
}

func (s *NoteService) Create(createNoteParams ports.CreateNoteParams) error {
	return s.noteRepository.Create(createNoteParams)
}

func (s *NoteService) GetAll() (ports.GetAllResponse, error) {
	return s.noteRepository.GetAll()
}
