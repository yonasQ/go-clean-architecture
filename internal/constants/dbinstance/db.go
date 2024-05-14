package dbinstance

import (
	"context"
	"project-structure-template/internal/constants/model/db"
	"project-structure-template/internal/constants/model/dto"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store interface {
	db.Querier
	GetAllUsers(ctx context.Context) ([]dto.User, error)
}

type DBInstance struct {
	*db.Queries
	Pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) Store {
	return &DBInstance{
		Pool:    pool,
		Queries: db.New(pool),
	}
}
