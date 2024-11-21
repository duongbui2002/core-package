package consumer

import (
	"context"
	"github.com/duongbui2002/core-package/core/messaging/types"
)

type ConsumerHandler interface {
	Handle(ctx context.Context, consumeContext types.MessageConsumeContext) error
}
