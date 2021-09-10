package transport

import (
	//lint:ignore ST1001 statically include native STTP types as root
	. "github.com/sttp/goapi/sttp"
)

// Function pointer types
type DispatcherFunction func(*DataSubscriber, []byte)
type MessageCallback func(*DataSubscriber, string)
type DataStartTimeCallback func(*DataSubscriber, Ticks)
type MetadataCallback func(*DataSubscriber, []byte)
type SubscriptionUpdatedCallback func(*DataSubscriber, *SignalIndexCache)
type NewMeasurementsCallback func(*DataSubscriber, []*Measurement)
type ConfigurationChangedCallback func(*DataSubscriber)
type ConnectionTerminatedCallback func(*DataSubscriber)

type DataSubscriber struct {
	subscriptionInfo SubscriptionInfo
}

func (ds *DataSubscriber) SetSubscriptionInfo(info SubscriptionInfo) {
	ds.subscriptionInfo = info
}
