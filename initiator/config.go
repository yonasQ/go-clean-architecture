package initiator

import (
	"context"
	"fmt"
	"project-structure-template/platform/logger"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig(name, path string, log logger.Logger) {
	viper.SetConfigName(name)
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(context.Background(), fmt.Sprintf("Failed to read config: %v", err))
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Panic(context.Background(), "Config file changed:", zap.String("file", e.Name))
	})
}
