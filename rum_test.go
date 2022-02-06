package rum

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestBootstrap(t *testing.T) {
	viper.SetDefault("app", "rum")
	Bootstrap("rum")
	fmt.Println(viper.GetString("app"))
}
