package postgres

import (
	gen "api-gateway/internal/adapters/repository/postgres/gen"
	"api-gateway/internal/core/domain"
	"api-gateway/internal/core/ports"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *APIGatewayRepository) Create(createNoteParams ports.CreateNoteParams) error {
	queries := gen.New(r.db)

	note_uuid := uuid.New().String()

	user_uuid_text := pgtype.Text{String: createNoteParams.UserUuid, Valid: true}

	params := gen.CreateNoteParams{
		Uuid:      note_uuid,
		Title:     createNoteParams.Title,
		Content:   createNoteParams.Content,
		CreatedBy: user_uuid_text,
	}

	_, err := queries.CreateNote(r.ctx, params)

	if err != nil {
		return err
	}

	//log.Println("New note created :", title)
	return nil
}

func (r *APIGatewayRepository) GetAll() (ports.GetAllResponse, error) {
	queries := gen.New(r.db)

	res, err := queries.GetAllNotes(r.ctx)

	if err != nil {
		return ports.GetAllResponse{}, err
	}

	var notes []*domain.Note

	for _, note := range res {
		notes = append(notes, &domain.Note{
			Title:   note.Title,
			Content: note.Content,
		})
	}

	return ports.GetAllResponse{
		Notes: notes,
	}, nil
}
