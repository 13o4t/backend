package config

import (
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
	"github.com/spf13/viper"
)

func Local(prefix, name string) (string, error) {
	if ex, err := os.Executable(); err == nil {
		dir := filepath.Dir(ex)
		viper.AddConfigPath(dir)
	}

	viper.AddConfigPath(configdir.LocalConfig(prefix))
	viper.SetConfigName(name)

	if err := viper.ReadInConfig(); err != nil {
		return "", err
	}
	return viper.ConfigFileUsed(), nil
}

func LocalSpecified(configFile string) error {
	viper.SetConfigFile(configFile)
	return viper.ReadInConfig()
}
