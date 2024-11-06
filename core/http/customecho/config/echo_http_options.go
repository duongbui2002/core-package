package config

import (
	"fmt"
	"github.com/duongbui2002/core-package/core/config"
	"github.com/duongbui2002/core-package/core/config/environment"
	typeMapper "github.com/duongbui2002/core-package/core/reflection/typemapper"
	"github.com/iancoleman/strcase"
	"net/url"
)

var optionName = strcase.ToLowerCamel(typeMapper.GetGenericTypeNameByT[EchoHttpOptions]())

type EchoHttpOptions struct {
	Port                string   `mapstructure:"port"                validate:"required" env:"TcpPort"`
	Development         bool     `mapstructure:"development"                             env:"Development"`
	BasePath            string   `mapstructure:"basePath"            validate:"required" env:"BasePath"`
	DebugErrorsResponse bool     `mapstructure:"debugErrorsResponse"                     env:"DebugErrorsResponse"`
	IgnoreLogUrls       []string `mapstructure:"ignoreLogUrls"`
	Timeout             int      `mapstructure:"timeout"                                 env:"Timeout"`
	Host                string   `mapstructure:"host"                                    env:"Host"`
	Name                string   `mapstructure:"name"                                    env:"ShortTypeName"`
}

func (c *EchoHttpOptions) Address() string {
	return fmt.Sprintf("%s%s", c.Host, c.Port)
}

func (c *EchoHttpOptions) BasePathAddress() string {
	path, err := url.JoinPath(c.Address(), c.BasePath)
	if err != nil {
		return ""
	}
	return path
}
func ProvideConfig(environment environment.Environment) (*EchoHttpOptions, error) {
	return config.BindConfigKey[*EchoHttpOptions](optionName, environment)
}
