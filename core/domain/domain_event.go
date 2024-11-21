package domain

import (
	"github.com/duongbui2002/core-package/core/events"
	uuid "github.com/satori/go.uuid"
)

type IDomainEvent interface {
	events.IEvent
	GetAggregateId() uuid.UUID
	GetAggregateSequenceNumber() int64
	WithAggregate(aggregateId uuid.UUID, aggregateSequenceNumber int64) *DomainEvent
}

type DomainEvent struct {
	*events.Event
	AggregateId             uuid.UUID `json:"aggregate_id"`
	AggregateSequenceNumber int64     `json:"aggregate_sequence_number"`
}

func (d *DomainEvent) GetAggregateId() uuid.UUID {
	return d.AggregateId
}

func (d *DomainEvent) GetAggregateSequenceNumber() int64 {
	return d.AggregateSequenceNumber
}

func (d *DomainEvent) WithAggregate(
	aggregateId uuid.UUID,
	aggregateSequenceNumber int64,
) *DomainEvent {
	d.AggregateId = aggregateId
	d.AggregateSequenceNumber = aggregateSequenceNumber

	return d
}
