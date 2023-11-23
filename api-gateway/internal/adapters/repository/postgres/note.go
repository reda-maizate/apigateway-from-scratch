package postgres

import (
	gen "api-gateway/internal/adapters/repository/postgres/gen"
	"api-gateway/internal/core/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (p *APIGatewayRepository) Create(title, content, userUuid string) error {
	queries := gen.New(p.db)

	note_uuid := uuid.New().String()

	user_uuid_text := pgtype.Text{String: userUuid, Valid: true}

	params := gen.CreateNoteParams{
		Uuid:      note_uuid,
		Title:     title,
		Content:   content,
		CreatedBy: user_uuid_text,
	}

	_, err := queries.CreateNote(p.ctx, params)

	if err != nil {
		return err
	}

	//log.Println("New note created :", title)
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
