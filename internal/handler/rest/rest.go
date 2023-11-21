package rest

import "github.com/gin-gonic/gin"

type User interface {
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}
