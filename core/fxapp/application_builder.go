package fxapp

import (
	"github.com/duongbui2002/core-package/core/config/environment"
	"github.com/duongbui2002/core-package/core/fxapp/contracts"
	"github.com/duongbui2002/core-package/core/logger"
	loggerConfig "github.com/duongbui2002/core-package/core/logger/config"
	"github.com/duongbui2002/core-package/core/logger/logrous"
	"github.com/duongbui2002/core-package/core/logger/models"

	"go.uber.org/fx"
)

type applicationBuilder struct {
	provides    []interface{}
	decorates   []interface{}
	options     []fx.Option
	logger      logger.Logger
	environment environment.Environment
}

func NewApplicationBuilder(environments ...environment.Environment) contracts.ApplicationBuilder {
	env := environment.ConfigAppEnv(environments...)

	var logger logger.Logger
	logoption, err := loggerConfig.ProvideLogConfig(env)
	if err != nil || logoption == nil {
		logger = logrous.NewLogrusLogger(logoption, env)
	} else if logoption.LogType == models.Logrus {
		logger = logrous.NewLogrusLogger(logoption, env)
	} else {
		logger = logrous.NewLogrusLogger(logoption, env)
	}

	return &applicationBuilder{logger: logger, environment: env}

}

func (a *applicationBuilder) ProvideModule(module fx.Option) {
	a.options = append(a.options, module)
}

func (a *applicationBuilder) Provide(constructors ...interface{}) {
	a.provides = append(a.provides, constructors...)
}

func (a *applicationBuilder) Decorate(constructors ...interface{}) {
	a.decorates = append(a.decorates, constructors...)
}

func (a *applicationBuilder) Build() contracts.Application {
	app := NewApplication(a.provides, a.decorates, a.options, a.logger, a.environment)

	return app
}

func (a *applicationBuilder) GetProvides() []interface{} {
	return a.provides
}

func (a *applicationBuilder) GetDecorates() []interface{} {
	return a.decorates
}

func (a *applicationBuilder) Options() []fx.Option {
	return a.options
}

func (a *applicationBuilder) Logger() logger.Logger {
	return a.logger
}

func (a *applicationBuilder) Environment() environment.Environment {
	return a.environment
}
