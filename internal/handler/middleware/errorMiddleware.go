package middleware

import (
	"fmt"
	"net/http"
	"project-structure-template/internal/constants"
	"project-structure-template/internal/constants/errors"
	"project-structure-template/internal/constants/model"
	"project-structure-template/platform/logger"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/joomcode/errorx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func ErrorHandler(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			errStatusCode, r := CastErrorResponse(c, log)
			constants.ErrorResponse(c, errStatusCode, r)
			return
		}
		// else {
		// 	log.Error(c, "somewhere in the handlers code missing ctx.Error(err)")
		// 	constants.ErrorResponse(
		// 		c,
		// 		http.StatusInternalServerError,
		// 		&[]model.Response{
		// 			{
		// 				OK: false,
		// 				Error: &model.ErrorResponse{
		// 					Code:    http.StatusInternalServerError,
		// 					Message: "Unknown server error",
		// 				},
		// 			},
		// 		},
		// 	)
		// 	return
		// }
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

func CastErrorResponse(c *gin.Context, log logger.Logger) (int, *[]model.Response) {
	debugMode := viper.GetBool("debug")
	errStatusCode := http.StatusInternalServerError
	modelResponse := []model.Response{}
	response := model.ErrorResponse{}

	for i, e := range c.Errors {
		err := e.Unwrap()
		er := errorx.Cast(err)
		if er == nil {
			log.Error(c, "unknown errorx type error", zap.Error(err))
			response = model.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Unknown server error",
			}
			if debugMode {
				response.StackTrace = fmt.Sprintf("%+v", errorx.EnsureStackTrace(err))
			}
			if len(c.Errors) == 1 {
				errStatusCode = response.Code
			}
			modelResponse = append(
				modelResponse,
				model.Response{
					OK:    false,
					Error: &response,
				},
			)
			if len(c.Errors) == (i + 1) {
				return errStatusCode, &modelResponse
			}
			continue
		}

		code, ok := errors.ErrorMap[er.Type()]
		if !ok {
			log.Error(c, "unknown errorx type error", zap.Error(er))
			response = model.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Unknown server error",
			}
			if debugMode {
				response.Description = fmt.Sprintf("Error: %v", er)
				response.StackTrace = fmt.Sprintf("%+v", errorx.EnsureStackTrace(err))
			}
			if len(c.Errors) == 1 {
				errStatusCode = response.Code
			}
			modelResponse = append(
				modelResponse,
				model.Response{
					OK:    false,
					Error: &response,
				},
			)

		} else {
			response = model.ErrorResponse{
				Code:       code,
				Message:    er.Message(),
				FieldError: ErrorFields(er.Cause()),
			}
			if debugMode {
				response.Description = fmt.Sprintf("Error: %v", er)
				response.StackTrace = fmt.Sprintf("%+v", errorx.EnsureStackTrace(err))
			}

			if len(c.Errors) == 1 {
				errStatusCode = response.Code
			}

			modelResponse = append(
				modelResponse,
				model.Response{
					OK:    false,
					Error: &response,
				},
			)
		}
	}

	return errStatusCode, &modelResponse
}
