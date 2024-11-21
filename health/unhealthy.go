package health

import (
	"context"
	"github.com/duongbui2002/core-package/health/contracts"
)

type UnhealthyHealthService struct{}

func NewUnhealthyHealthService() UnhealthyHealthService {
	return UnhealthyHealthService{}
}

func (service UnhealthyHealthService) CheckHealth(
	context.Context,
) contracts.Check {
	return contracts.Check{
		"postgres": contracts.Status{Status: contracts.StatusDown},
		"redis":    contracts.Status{Status: contracts.StatusDown},
	}
}
