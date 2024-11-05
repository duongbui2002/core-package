package config

import (
	"github.com/duongbuidinh600/core-package/core/config"
	"github.com/duongbuidinh600/core-package/core/config/environment"
	"github.com/duongbuidinh600/core-package/core/logger/models"
	typeMapper "github.com/duongbuidinh600/core-package/core/reflection/typemapper"
	"github.com/iancoleman/strcase"
)

var optionName = strcase.ToLowerCamel(typeMapper.GetGenericTypeNameByT[LogOptions]())

type LogOptions struct {
	LogLevel      string         `mapstructure:"level"`
	LogType       models.LogType `mapstructure:"logType"`
	CallerEnabled bool           `mapstructure:"callerEnabled"`
	EnableTracing bool           `mapstructure:"enableTracing" default:"true"`
}

func ProvideLogConfig(env environment.Environment) (*LogOptions, error) {
	return config.BindConfigKey[*LogOptions](optionName, env)
}
