package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// 目前只支持yaml配置文件

func InitConfig(configDir string, ConfigName string) {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetConfigName(ConfigName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed read config file, error: %s \n", err)
		os.Exit(1)
	}
	_RUNONCE.Do(_loadStructConfig)
}
