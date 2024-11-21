package elasticsearch

import (
	"github.com/duongbui2002/core-package/config"
	"github.com/duongbui2002/core-package/config/environment"
	typeMapper "github.com/duongbui2002/core-package/reflection/typemapper"
	"github.com/iancoleman/strcase"
)

var optionName = strcase.ToLowerCamel(typeMapper.GetGenericTypeNameByT[ElasticOptions]())

type ElasticOptions struct {
	URL string `mapstructure:"url"`
}

func provideConfig(environment environment.Environment) (*ElasticOptions, error) {
	return config.BindConfigKey[*ElasticOptions](optionName, environment)
}
