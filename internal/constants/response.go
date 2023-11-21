package constants

import (
	"project-structure-template/internal/constants/model"

	"github.com/gin-gonic/gin"
)

func SuccessResponse(ctx *gin.Context, statusCode int, data interface{}, metaData *model.MetaData) {
	ctx.JSON(
		statusCode,
		model.Response{
			OK:       true,
			MetaData: metaData,
			Data:     data,
		},
	)
}

func ErrorResponse(ctx *gin.Context, err *model.ErrorResponse) {
	ctx.AbortWithStatusJSON(err.Code, model.Response{
		OK:    false,
		Error: err,
	})
}
