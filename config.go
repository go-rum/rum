package rum

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	supportiveConfigType = "toml"
)

func SetSupportiveConfigTypes(configType string) {
	supportiveConfigType = configType
}

func LoadConfig(name string) error {
	viper.SetConfigName(name) // name of config file (without extension)

	viper.SetConfigType(supportiveConfigType) // REQUIRED if the config file does not have the extension in the name

	viper.AddConfigPath(fmt.Sprintf("/etc/%s/", name))  // path to look for the config file in
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", name)) // call multiple times to add many search paths
	viper.AddConfigPath(".")                            // optionally look for config in the working directory

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return fmt.Errorf("viper read config: %w", err)
	}

	return nil
}
