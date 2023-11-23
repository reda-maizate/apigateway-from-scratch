package services

import (
	"api-gateway/internal/core/domain"
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

func (s *NoteService) Create(title, content, userUuid string) error {
	return s.noteRepository.Create(title, content, userUuid)
}

func (s *NoteService) GetAll() ([]*domain.Note, error) {
	return s.noteRepository.GetAll()
}
