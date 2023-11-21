package module

import (
	"context"
	"project-structure-template/internal/constants/model/dto"
)

type User interface {
	Create(ctx context.Context, param dto.RegisterUser) (*dto.User, error)
	Update(ctx context.Context, id string, param dto.UpdateUser) (*dto.User, error)
	GetAll(ctx context.Context) ([]dto.User, error)
	Get(ctx context.Context, id string) (*dto.User, error)
	DeleteUser(ctx context.Context, userId string) error
}
