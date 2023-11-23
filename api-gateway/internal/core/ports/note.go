package ports

import "api-gateway/internal/core/domain"

type CreateNoteParams struct {
	Title    string
	Content  string
	UserUuid string
}

type GetAllResponse struct {
	Notes []*domain.Note
}

type NoteService interface {
	Create(CreateNoteParams) error
	GetAll() (GetAllResponse, error)
}

type NoteRepository interface {
	Create(CreateNoteParams) error
	GetAll() (GetAllResponse, error)
}
