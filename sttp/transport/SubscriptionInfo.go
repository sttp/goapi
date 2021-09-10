package transport

type SubscriptionInfo struct {
	FilterExpression string

	// Down-sampling properties
	Throttled       bool
	PublishInterval float64

	// UDP channel properties
	UdpDataChannel       bool
	DataChannelLocalPort uint16

	// Compact measurement properties
	IncludeTime              bool
	LagTime                  float64
	LeadTime                 float64
	UseLocalClockAsRealTime  bool
	UseMillisecondResolution bool
	RequestNaNValueFilter    bool

	// Temporal playback properties
	StartTime            string
	StopTime             string
	ConstraintParameters string
	ProcessingInterval   int32

	ExtraConnectionStringParameters string
}
