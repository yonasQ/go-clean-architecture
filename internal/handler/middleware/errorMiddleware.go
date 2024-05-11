package middleware

import (
	"fmt"
	"net/http"
	"project-structure-template/internal/constants"
	"project-structure-template/internal/constants/errors"
	"project-structure-template/internal/constants/model"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/joomcode/errorx"
	"github.com/spf13/viper"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			e := c.Errors[0] // FIXME: how to handle multiple errors?
			err := e.Unwrap()

			constants.ErrorResponse(c, CastErrorResponse(err))
			return
		}
	}
}

func ErrorFields(err error) []model.FieldError {
	var errs []model.FieldError

	if data, ok := err.(validation.Errors); ok {
		for i, v := range data {
			errs = append(
				errs, model.FieldError{
					Name:        i,
					Description: v.Error(),
				},
			)
		}

		return errs
	}

	return nil
}

func CastErrorResponse(err error) *model.ErrorResponse {
	debugMode := viper.GetBool("debug")
	er := errorx.Cast(err)
	if er == nil {
		return &model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Unknown server error",
		}
	}

	response := model.ErrorResponse{}
	code, ok := errors.ErrorMap[er.Type()]
	if !ok {
		response = model.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Unknown server error",
		}
	} else {
		response = model.ErrorResponse{
			Code:       code,
			Message:    er.Message(),
			FieldError: ErrorFields(er.Cause()),
		}
	}

	if debugMode {
		response.Description = fmt.Sprintf("Error: %v", er)
		response.StackTrace = fmt.Sprintf("%+v", errorx.EnsureStackTrace(err))
	}

	return &response
}
