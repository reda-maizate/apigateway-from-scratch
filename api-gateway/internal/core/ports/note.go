package ports

import "api-gateway/internal/core/domain"

type NoteService interface {
	Create(title, content, userUuid string) error
	GetAll() ([]*domain.Note, error)
}

type NoteRepository interface {
	Create(title, content, userUuid string) error
	GetAll() ([]*domain.Note, error)
}
