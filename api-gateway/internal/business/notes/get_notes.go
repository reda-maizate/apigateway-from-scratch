package notes

import (
	database "api-gateway/internal/db"
	"context"
)

func (b *NotesBusiness) GetAllNotes(ctx context.Context) ([]*database.Note, error) {
	res, err := b.queries.GetAllNotes(ctx)

	if err != nil {
		return nil, err
	}

	var notes []*database.Note

	for _, note := range res {
		notes = append(notes, &database.Note{
			Title:   note.Title,
			Content: note.Content,
		})
	}

	return notes, nil
}
