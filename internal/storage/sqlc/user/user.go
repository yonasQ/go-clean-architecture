package user

import (
	"context"
	goerror "errors"
	"project-structure-template/internal/constants/dbinstance"
	"project-structure-template/internal/constants/errors"
	"project-structure-template/internal/constants/model/db"
	"project-structure-template/internal/constants/model/dto"
	"project-structure-template/internal/storage"
	"project-structure-template/platform/logger"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type user struct {
	db  dbinstance.DBInstance
	log logger.Logger
}

func Init(db dbinstance.DBInstance, log logger.Logger) storage.User {
	return &user{
		db:  db,
		log: log,
	}
}
func (u *user) Create(ctx context.Context, param dto.RegisterUser) (*dto.User, error) {
	user, err := u.db.CreateUser(ctx, db.CreateUserParams{
		Email:      param.Email,
		FirstName:  param.FirstName,
		MiddleName: param.MiddleName,
		LastName:   param.LastName,
	})
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not create user")
		u.log.Error(ctx, "unable to create user", zap.Error(err), zap.Any("user", param))
		return nil, err
	}
	return &dto.User{
		ID:         user.ID,
		Email:      user.Email,
		FirstName:  user.FirstName,
		MiddleName: user.MiddleName,
		LastName:   user.LastName,
		CreatedAt:  user.CreatedAt,
	}, nil
}

func (u *user) Update(ctx context.Context, id uuid.UUID, param dto.UpdateUser) (*dto.User, error) {
	user, err := u.db.UpdateUser(ctx, db.UpdateUserParams{
		FirstName:  param.FirstName,
		MiddleName: param.MiddleName,
		LastName:   param.LastName,
		ID:         id,
	})
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not update user")
		u.log.Error(ctx, "unable to update user", zap.Error(err), zap.Any("user", param))
		return nil, err
	}
	return &dto.User{
		ID:         user.ID,
		Email:      user.Email,
		FirstName:  user.FirstName,
		MiddleName: user.MiddleName,
		LastName:   user.LastName,
		CreatedAt:  user.CreatedAt,
	}, nil
}

func (u *user) Get(ctx context.Context, id uuid.UUID) (*dto.User, error) {
	user, err := u.db.GetUser(ctx, id)
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not read user")
		u.log.Error(ctx, "unable to get user", zap.Error(err))
		return nil, err
	}
	return &dto.User{
		ID:         user.ID,
		Email:      user.Email,
		FirstName:  user.FirstName,
		MiddleName: user.MiddleName,
		LastName:   user.LastName,
		CreatedAt:  user.CreatedAt,
		Status:     string(user.Status),
	}, nil
}

func (u *user) GetAll(ctx context.Context) ([]dto.User, error) {
	users, err := u.db.GetAllUsers(ctx)
	if err != nil {
		err = errors.ErrWriteError.Wrap(err, "could not read users")
		u.log.Error(ctx, "unable to get users", zap.Error(err))
		return nil, err
	}
	return users, nil
}

func (u *user) IsUserExists(ctx context.Context, param dto.RegisterUser) (bool, error) {
	count, err := u.db.UserByEmailExists(ctx, param.Email)
	if err != nil {
		err := errors.ErrReadError.Wrap(err, "could not read user")
		u.log.Error(ctx, "unable to read the user", zap.Error(err), zap.Any("user-email", param.Email))
		return false, err
	}

	if count.(int64) > 0 {
		return true, nil
	}
	return false, nil
}

func (u *user) DeleteUser(ctx context.Context, userId uuid.UUID) error {
	_, err := u.db.DeleteUser(ctx, userId)
	if err != nil {
		if err.Error() == goerror.New("no rows in result set").Error() {
			err := errors.ErrNoRecordFound.Wrap(err, "no record of user found")
			u.log.Info(ctx, "User with this id not found", zap.Error(err), zap.String("user-id", userId.String()))
			return err
		}
		err = errors.ErrWriteError.Wrap(err, "error deleting the user")
		u.log.Error(ctx, "unable to delete user data", zap.Error(err), zap.String("user-id", userId.String()))
		return err
	}
	return nil
}
