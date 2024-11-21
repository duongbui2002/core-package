package elasticsearch

import "go.uber.org/fx"

var Module = fx.Module("elasticfx",
	fx.Provide(provideConfig),
	fx.Provide(NewElasticClient),
)
