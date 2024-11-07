package client

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"client-module",
	fx.Provide(NewHttpClient),
)
