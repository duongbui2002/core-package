package domain

import (
	"github.com/duongbui2002/core-package/core/metadata"
)

type EventEnvelope struct {
	EventData interface{}
	Metadata  metadata.Metadata
}
