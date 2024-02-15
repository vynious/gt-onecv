package db

import (
	"github.com/jackc/pgx/v5"
	"github.com/vynious/gt-onecv/db/sqlc"
)

type Repository struct {
	Queries *sqlc.Queries
	DB      *pgx.Conn
}

func SpawnRepository(db *pgx.Conn) *Repository {
	return &Repository{
		DB:      db,
		Queries: sqlc.New(db),
	}
}
