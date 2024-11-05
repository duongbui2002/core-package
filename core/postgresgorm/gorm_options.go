package postgresgorm

import (
	"fmt"
	"github.com/duongbuidinh600/core-package/core/config"
	"github.com/duongbuidinh600/core-package/core/config/environment"
	typeMapper "github.com/duongbuidinh600/core-package/core/reflection/typemapper"
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
		projectRootDir := environment.GetProjectRootWorkingDirectory()
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

func provideConfig(environment environment.Environment) (*GormOptions, error) {
	return config.BindConfigKey[*GormOptions](optionName, environment)
}
