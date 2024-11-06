package contracts

import (
	"context"
	"github.com/duongbui2002/core-package/core/config/environment"
	"github.com/duongbui2002/core-package/core/logger"
	"go.uber.org/fx"
)

type Application interface {
	Container
	RegisterHook(function interface{})
	Run()
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Wait() <-chan fx.ShutdownSignal
	Logger() logger.Logger
	Environment() environment.Environment
}
