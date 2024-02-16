package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vynious/gt-onecv/db/sqlc"
)

type Repository struct {
	Queries *sqlc.Queries
	DB      *pgxpool.Pool
}

func SpawnRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		DB:      db,
		Queries: sqlc.New(db),
	}
}
