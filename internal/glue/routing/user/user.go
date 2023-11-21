package user

import (
	"net/http"
	"project-structure-template/internal/glue/routing"
	"project-structure-template/internal/handler/rest"

	"github.com/gin-gonic/gin"
)

func InitRoute(grp *gin.RouterGroup, user rest.User) {
	users := grp.Group("users")
	usersRoutes := []routing.Router{
		{
			Method:  http.MethodPost,
			Handler: user.CreateUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/:id",
			Handler: user.UpdateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/:id",
			Handler: user.GetUser,
		},
		{
			Method:  http.MethodGet,
			Handler: user.GetUsers,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/:id",
			Handler: user.DeleteUser,
		},
	}
	routing.RegisterRoutes(users, usersRoutes)
}
