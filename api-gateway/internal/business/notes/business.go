package notes

import (
	core_business "api-gateway/internal/business/core"
	database "api-gateway/internal/db"
	"github.com/jackc/pgx/v5"
)

type NotesBusiness struct {
	db      *pgx.Conn
	queries *database.Queries
}

type NotesBusinessConfig struct{}

func NewNotesBusiness(coreConfig *core_business.CoreBusinessConfig, config *NotesBusinessConfig) *NotesBusiness {
	return &NotesBusiness{
		db:      coreConfig.DB,
		queries: coreConfig.Queries,
	}
}
