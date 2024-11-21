package postgresgorm

import (
	"fmt"
	"github.com/duongbui2002/core-package/config"
	environment2 "github.com/duongbui2002/core-package/config/environment"
	typeMapper "github.com/duongbui2002/core-package/reflection/typemapper"
	"github.com/iancoleman/strcase"
	"path/filepath"
)

var optionName = strcase.ToLowerCamel(typeMapper.GetGenericTypeNameByT[GormOptions]())

type GormOptions struct {
	UseInMemory   bool   `mapstructure:"useInMemory"`
	UseSQLLite    bool   `mapstructure:"useSqlLite"`
	Host          string `mapstructure:"host"`
	Port          int    `mapstructure:"port"`
	User          string `mapstructure:"user"`
	DBName        string `mapstructure:"dbName"`
	SSLMode       bool   `mapstructure:"sslMode"`
	Password      string `mapstructure:"password"`
	EnableTracing bool   `mapstructure:"enableTracing" default:"true"`
}

func (h *GormOptions) Dns() string {
	if h.UseInMemory {
		return ""
	}

	if h.UseSQLLite {
		projectRootDir := environment2.GetProjectRootWorkingDirectory()
		dbFilePath := filepath.Join(projectRootDir, fmt.Sprintf("%s.db", h.DBName))

		return dbFilePath
	}

	datasource := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		h.User,
		h.Password,
		h.Host,
		h.Port,
		h.DBName,
	)

	return datasource
}

func provideConfig(environment environment2.Environment) (*GormOptions, error) {
	return config.BindConfigKey[*GormOptions](optionName, environment)
}
