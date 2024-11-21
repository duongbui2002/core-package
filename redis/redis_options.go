package redis

import (
	"github.com/duongbui2002/core-package/config"
	"github.com/duongbui2002/core-package/config/environment"
	typeMapper "github.com/duongbui2002/core-package/reflection/typemapper"
	"github.com/iancoleman/strcase"
)

var optionName = strcase.ToLowerCamel(typeMapper.GetGenericTypeNameByT[RedisOptions]())

type RedisOptions struct {
	Host          string `mapstructure:"host"`
	Port          int    `mapstructure:"port"`
	Password      string `mapstructure:"password"`
	Database      int    `mapstructure:"database"`
	PoolSize      int    `mapstructure:"poolSize"`
	EnableTracing bool   `mapstructure:"enableTracing" default:"true"`
}

func provideConfig(environment environment.Environment) (*RedisOptions, error) {
	return config.BindConfigKey[*RedisOptions](optionName, environment)
}
