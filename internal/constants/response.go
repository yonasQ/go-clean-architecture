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

func ErrorResponse(ctx *gin.Context, errStatusCode int, r *[]model.Response, multiple bool) {
	if !multiple {
		r := *r
		ctx.AbortWithStatusJSON(errStatusCode, r[0])
		return
	}
	ctx.AbortWithStatusJSON(errStatusCode, r)
}
