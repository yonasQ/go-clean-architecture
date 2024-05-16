package user

import (
	"context"
	"net/http"
	"project-structure-template/internal/constants"
	"project-structure-template/internal/constants/errors"
	"project-structure-template/internal/constants/model/dto"
	"project-structure-template/internal/handler/rest"
	"project-structure-template/internal/module"
	"project-structure-template/platform/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type user struct {
	logger         logger.Logger
	userModule     module.User
	contextTimeout time.Duration
}

func Init(logger logger.Logger, userModule module.User, contextTimeout time.Duration) rest.User {
	return &user{
		userModule:     userModule,
		logger:         logger,
		contextTimeout: contextTimeout,
	}
}

// CreateUser creates a new user.
//
//	@Summary		Create a new user.
//	@Description	This endpoint is used to create a new user by providing the necessary details in the request body.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.RegisterUser	true	"User details in JSON format"
//	@Success		201		{object}	dto.User			"Successfully created user"
//	@Failure		400		{object}	model.Response		"Bad request, check the error response for details"
//	@Router			/users [post]
func (u *user) CreateUser(ctx *gin.Context) {
	cntx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user := dto.RegisterUser{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		u.logger.Error(ctx, "unable to bind user data", zap.Error(err))
		_ = ctx.Error(err)
		return
	}

	createdUser, err := u.userModule.Create(cntx, user)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	constants.SuccessResponse(ctx, http.StatusCreated, createdUser, nil)
}

// UpdateUser updates a user by ID.
//
//	@Summary		Update an existing user.
//	@Description	This endpoint is used to update an existing user identified by the provided User ID.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string			true	"User ID"
//	@Param			user	body		dto.UpdateUser	true	"Updated user details in JSON format"
//	@Success		200		{object}	dto.User		"Successfully updated user"
//	@Failure		400		{object}	model.Response	"Bad request, check the error response for details"
//	@Router			/users/{id} [patch]
func (u *user) UpdateUser(ctx *gin.Context) {
	cntx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user := dto.UpdateUser{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		u.logger.Error(ctx, "unable to bind user data", zap.Error(err))
		_ = ctx.Error(err)
		return
	}

	updatedUser, err := u.userModule.Update(cntx, ctx.Param("id"), user)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	constants.SuccessResponse(ctx, http.StatusOK, updatedUser, nil)
}

// GetUser gets a user by ID.
//
//	@Summary		Get user by ID.
//	@Description	This endpoint is used to retrieve information about a user identified by the provided User ID.
//	@Tags			user
//	@Produce		json
//	@Param			id	path		string			true	"User ID"
//	@Success		200	{object}	dto.User		"Successfully retrieved user"
//	@Failure		400	{object}	model.Response	"Bad request, check the error response for details"
//	@Router			/users/{id} [get]
func (u *user) GetUser(ctx *gin.Context) {
	cntx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err := u.userModule.Get(cntx, ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	constants.SuccessResponse(ctx, http.StatusCreated, user, nil)
}

// GetUser gets a list of users.
//
//	@Summary		Get users.
//	@Description	This endpoint is used to retrieve a list of users.
//	@Tags			user
//	@Produce		json
//	@Success		200	{object}	dto.User		"Successfully retrieved users"
//	@Failure		400	{object}	model.Response	"Bad request, check the error response for details"
//	@Router			/users [get]
func (u *user) GetUsers(ctx *gin.Context) {
	cntx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	users, err := u.userModule.GetAll(cntx)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	constants.SuccessResponse(ctx, http.StatusCreated, users, nil)
}

// DeleteUser is used to delete a user.
//
//	@Summary		Delete user.
//	@Description	This function deletes a user if the user is available.
//	@Tags			users
//	@Produce		json
//	@Param			id	path	string			true	"User ID"
//	@Success		200	string	string			"Successfully deleted the user"
//	@Success		404	string	model.Response	"User not found"
//	@Success		400	string	model.Response	"Invalid user ID"
//	@Router			/users/{id} [delete]
func (u *user) DeleteUser(ctx *gin.Context) {
	cntx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	err := u.userModule.DeleteUser(cntx, ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	constants.SuccessResponse(ctx, http.StatusOK, "User deleted successfully", nil)
}
