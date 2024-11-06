package fxapp

import (
	"github.com/duongbui2002/core-package/core/config"
	"github.com/duongbui2002/core-package/core/logger/external/fxlog"
	"github.com/duongbui2002/core-package/core/logger/logrous"
	"time"

	"go.uber.org/fx"
)

func CreateFxApp(
	app *application,
) *fx.App {
	var opts []fx.Option

	opts = append(opts, fx.Provide(app.provides...))

	opts = append(opts, fx.Decorate(app.decorates...))

	opts = append(opts, fx.Invoke(app.invokes...))

	app.options = append(app.options, opts...)

	AppModule := fx.Module("fxapp",
		app.options...,
	)

	var logModule fx.Option

	logModule = logrous.ModuleFunc(app.logger)

	duration := 30 * time.Second

	// build phase of container will do in this stage, containing provides and invokes but app not started yet and will be started in the future with `fxApp.Register`
	fxApp := fx.New(
		fx.StartTimeout(duration),
		config.ModuleFunc(app.environment),
		logModule,
		fxlog.FxLogger,
		fx.ErrorHook(NewFxErrorHandler(app.logger)),
		AppModule,
	)

	return fxApp
}
