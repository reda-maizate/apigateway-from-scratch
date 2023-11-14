package services

import (
	"api-gateway/internal/core/domain"
	"api-gateway/internal/core/ports"
)

type NoteService struct {
	repo ports.NoteRepository
}

func NewNoteService(repo ports.NoteRepository) *NoteService {
	return &NoteService{
		repo: repo,
	}
}

func (s *NoteService) Create(title, content string) error {
	return s.repo.Create(title, content)
}

func (s *NoteService) GetAll() ([]*domain.Note, error) {
	return s.repo.GetAll()
}