package user

import (
	"context"
	"project-structure-template/internal/constants/errors"
	"project-structure-template/internal/constants/model/dto"
	"project-structure-template/internal/module"
	"project-structure-template/internal/storage"
	"project-structure-template/platform/logger"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type user struct {
	userPersistant storage.User
	log            logger.Logger
}

func Init(log logger.Logger, userPersistant storage.User) module.User {
	return &user{
		userPersistant: userPersistant,
		log:            log,
	}
}

func (u *user) Create(ctx context.Context, param dto.RegisterUser) (*dto.User, error) {
	if err := param.Validate(); err != nil {
		err = errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		u.log.Error(ctx, "validation failed", zap.Error(err), zap.Any("input", param))
		return nil, err
	}

	isExist, err := u.userPersistant.IsUserExists(ctx, param)
	if err != nil {
		return nil, err
	}

	if isExist {
		err = errors.ErrDataExists.New("user with this email already exists")
		u.log.Error(ctx, "duplicated data", zap.String("user-email", param.Email))
		return nil, err
	}

	return u.userPersistant.Create(ctx, param)
}

func (u *user) DeleteUser(ctx context.Context, id string) error {
	userId, err := uuid.Parse(id)
	if err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "invalid user id")
		u.log.Error(ctx, "parsing user id failed", zap.Error(err), zap.String("user-id", id))
		return err
	}

	return u.userPersistant.DeleteUser(ctx, userId)
}

func (u *user) Update(ctx context.Context, id string, param dto.UpdateUser) (*dto.User, error) {
	if err := param.Validate(); err != nil {
		err = errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		u.log.Error(ctx, "validation failed", zap.Error(err), zap.Any("input", param))
		return nil, err
	}

	uuidID, err := uuid.Parse(id)
	if err != nil {
		err = errors.ErrInvalidUserInput.Wrap(err, "invalid user id")
		u.log.Error(ctx, "parsing user id failed", zap.Error(err))
		return nil, err
	}

	return u.userPersistant.Update(ctx, uuidID, param)
}

func (u *user) Get(ctx context.Context, id string) (*dto.User, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "invalid user id")
		u.log.Error(ctx, "parsing user id failed", zap.Error(err))
		return nil, err
	}

	return u.userPersistant.Get(ctx, uuidID)
}

func (u *user) GetAll(ctx context.Context) ([]dto.User, error) {
	return u.userPersistant.GetAll(ctx)
}
