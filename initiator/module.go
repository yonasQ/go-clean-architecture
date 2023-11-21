package initiator

import (
	"project-structure-template/internal/module"
	"project-structure-template/internal/module/user"
	"project-structure-template/platform/logger"
)

type Module struct {
	user module.User
}

func InitModule(persistence Persistence, log logger.Logger) Module {
	return Module{
		user: user.Init(log.Named("user-module"), persistence.user),
	}
}
