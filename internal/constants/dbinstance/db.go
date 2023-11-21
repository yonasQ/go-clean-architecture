package dbinstance

import (
	"project-structure-template/internal/constants/model/db"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DBInstance struct {
	*db.Queries
	Pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) DBInstance {
	return DBInstance{
		Pool:    pool,
		Queries: db.New(pool),
	}
}
