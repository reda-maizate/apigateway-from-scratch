package postgres

import (
	"api-gateway/internal/adapters/repository/postgres/gen"
	"api-gateway/internal/core/domain"
	"log"
)

func (p *APIGatewayRepository) Create(title, content string) error {
	queries := gen.New(p.db)

	params := gen.CreateNoteParams{
		Title:   title,
		Content: content,
	}

	_, err := queries.CreateNote(p.ctx, params)

	if err != nil {
		return err
	}

	log.Println("New note created :", title)
	return nil
}

func (p *APIGatewayRepository) GetAll() ([]*domain.Note, error) {
	queries := gen.New(p.db)

	res, err := queries.GetAllNotes(p.ctx)

	if err != nil {
		return nil, err
	}

	var notes []*domain.Note

	for _, note := range res {
		notes = append(notes, &domain.Note{
			Title:   note.Title,
			Content: note.Content,
		})
	}

	return notes, nil
}
