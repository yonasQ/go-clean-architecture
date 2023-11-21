package storage

import (
	"context"
	"project-structure-template/internal/constants/model/dto"

	"github.com/google/uuid"
)

type User interface {
	Create(ctx context.Context, param dto.RegisterUser) (*dto.User, error)
	Update(ctx context.Context, id uuid.UUID, param dto.UpdateUser) (*dto.User, error)
	IsUserExists(ctx context.Context, param dto.RegisterUser) (bool, error)
	GetAll(ctx context.Context) ([]dto.User, error)
	Get(ctx context.Context, id uuid.UUID) (*dto.User, error)
	DeleteUser(ctx context.Context, userId uuid.UUID) error
}
