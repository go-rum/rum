package rum

import (
	"fmt"
	"github.com/spf13/viper"
)

func Bootstrap(appName string) {
	err := LoadConfig(appName)
	if err != nil {
		panic(fmt.Errorf("load config: %w", err))
	}

	if viper.GetString("mysql.host") != "" {
		LoadGormDB()
	}
}
