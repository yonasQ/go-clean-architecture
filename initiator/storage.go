package initiator

import (
	"project-structure-template/internal/constants/dbinstance"
	"project-structure-template/internal/storage"
	"project-structure-template/internal/storage/sqlc/user"
	"project-structure-template/platform/logger"
)

type Persistence struct {
	user storage.User
}

func InitPersistence(db dbinstance.DBInstance, log logger.Logger) Persistence {
	return Persistence{
		user: user.Init(db, log.Named("user-persistence")),
	}
}
