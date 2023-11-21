package initiator

import (
	"project-structure-template/internal/handler/rest"
	"project-structure-template/internal/handler/rest/gin/user"
	"project-structure-template/platform/logger"
	"time"
)

type Handler struct {
	user rest.User
}

func InitHandler(module Module, log logger.Logger, timeout time.Duration) Handler {
	return Handler{
		user: user.Init(log.Named("user-handler"), module.user, timeout),
	}
}
