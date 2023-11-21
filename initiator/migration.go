package initiator

import (
	"context"
	"fmt"
	"project-structure-template/platform/logger"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
)

func InitiateMigration(path, conn string, log logger.Logger) *migrate.Migrate {
	conn = fmt.Sprintf("cockroach://%s", strings.Split(conn, "://")[1])
	m, err := migrate.New(fmt.Sprintf("file://%s", path), conn)
	if err != nil {
		log.Fatal(context.Background(), "could not create migrator", zap.Error(err))
	}
	return m
}

func UpMigration(m *migrate.Migrate, log logger.Logger) {
	err := m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(context.Background(), "could not migrate", zap.Error(err))
	}
}
