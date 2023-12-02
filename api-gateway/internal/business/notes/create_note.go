package notes

import (
	database "api-gateway/internal/db"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateNoteParams struct {
	Title    string
	Content  string
	UserUuid string
}

func (b *NotesBusiness) CreateNote(ctx context.Context, params *CreateNoteParams) error {
	note_uuid := uuid.New().String()
	user_uuid_text := pgtype.Text{String: params.UserUuid, Valid: true}

	_, err := b.queries.CreateNote(ctx, database.CreateNoteParams{
		Uuid:      note_uuid,
		Title:     params.Title,
		Content:   params.Content,
		CreatedBy: user_uuid_text,
	})

	if err != nil {
		return err
	}

	//log.Println("New note created :", title)
	return nil
}
