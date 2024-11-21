package bus

import (
	"github.com/duongbui2002/core-package/core/messaging/consumer"
	"github.com/duongbui2002/core-package/core/messaging/producer"
)

type Bus interface {
	producer.Producer
	consumer.BusControl
	consumer.ConsumerConnector
}
