package transport

import (
	"time"
)

type ErrorMessageCallback func(*DataSubscriber, string)
type ReconnectCallback func(*DataSubscriber)

type SubscriberConnector struct {
	errorMessageCallback ErrorMessageCallback
	reconnectCallback    ReconnectCallback

	hostname string
	port     uint16
	timer    *time.Timer

	maxRetries        int32
	retryInterval     int32
	maxRetryInterval  int32
	connectAttempt    int32
	connectionRefused bool
	autoReconnect     bool
	cancel            bool
}

const ConnectSuccess int = 1
const ConnectFailed int = 0
const ConnectCanceled int = -1

// Registers a callback to provide error messages each time
// the subscriber fails to connect during a connection sequence.
func (sc *SubscriberConnector) RegisterErrorMessageCallback(errorMessageCallback ErrorMessageCallback) {
	sc.errorMessageCallback = errorMessageCallback
}

// Registers a callback to notify after an automatic reconnection attempt has been made.
// This callback will be called whether the connection was successful or not, so it is
// recommended to check the connected state of the subscriber using the IsConnected() method.
func (sc *SubscriberConnector) RegisterReconnectCallback(reconnectCallback ReconnectCallback) {
	sc.reconnectCallback = reconnectCallback
}

func (sc *SubscriberConnector) Connect(subscriber *DataSubscriber, info SubscriptionInfo) int {
	if sc.cancel {
		return ConnectCanceled
	}

	subscriber.SetSubscriptionInfo(info)
	return sc.connect(subscriber, false)
}

func (sc *SubscriberConnector) connect(subscriber *DataSubscriber, autoReconnecting bool) int {
	return ConnectSuccess
}
