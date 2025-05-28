package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"strings"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			NewConfig,
		),
	)
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil

}
