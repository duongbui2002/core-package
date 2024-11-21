package redis

import (
	"context"
	"fmt"
	"github.com/duongbui2002/core-package/health/contracts"
	"github.com/duongbui2002/core-package/logger"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

var (
	Module = fx.Module(
		"redisfx",
		redisProviders,
		redisInvokes,
	) //nolint:gochecknoglobals

	redisProviders = fx.Options(fx.Provide(
		NewRedisClient,
		func(client *redis.Client) redis.UniversalClient {
			return client
		},
		fx.Annotate(
			NewRedisHealthChecker,
			fx.As(new(contracts.Health)),
			fx.ResultTags(fmt.Sprintf(`group:"%s"`, "healths")),
		),
		provideConfig,
	))

	redisInvokes = fx.Options(
		fx.Invoke(registerHooks),
	)
)

func registerHooks(lc fx.Lifecycle,
	client redis.UniversalClient,
	logger logger.Logger,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {

			return client.Ping(ctx).Err()

		},
		OnStop: func(ctx context.Context) error {
			if err := client.Close(); err != nil {
				logger.Errorf("error in closing redis: %v", err)
			} else {
				logger.Info("redis closed gracefully")
			}

			return nil
		},
	})
}
