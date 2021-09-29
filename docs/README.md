## Go STTP ([IEEE 2664](https://standards.ieee.org/project/2664.html)) Documentation

<img align="right" src="img/sttp.png">

### Streaming Telemetry Transport Protocol

[![Go Report Card](https://goreportcard.com/badge/github.com/sttp/goapi)](https://goreportcard.com/report/github.com/sttp/goapi)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/sttp/goapi)](https://pkg.go.dev/github.com/sttp/goapi)
[![Release](https://img.shields.io/github/release/sttp/goapi.svg?style=flat-square)](https://github.com/sttp/goapi/releases/latest)

## Example Usage
```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "time"

    "github.com/sttp/goapi/sttp"
    "github.com/sttp/goapi/sttp/format"
    "github.com/sttp/goapi/sttp/transport"
)

// TestSubscriber is a simple STTP data subscriber implementation.
type TestSubscriber struct {
    sttp.SubscriberBase // Provides default implementation
}

// NewTestSubscriber creates a new TestSubscriber.
func NewTestSubscriber() *TestSubscriber {
    subscriber := &TestSubscriber{}
    subscriber.SubscriberBase = sttp.NewSubscriberBase(subscriber)
    return subscriber
}

func main() {
    subscriber := NewTestSubscriber()
    subscription := subscriber.Subscription()

    subscriber.Hostname = "localhost"
    subscriber.Port = 7165
    subscriber.Version = 1

    subscription.FilterExpression = "FILTER TOP 10 ActiveMeasurements WHERE SignalType = 'FREQ'"

    subscriber.Connect()
    defer subscriber.Dispose()

    reader := bufio.NewReader(os.Stdin)
    reader.ReadRune()
}

var lastMessageDisplay time.Time

// ReceivedNewMeasurements handles reception of new measurements.
func (sub *TestSubscriber) ReceivedNewMeasurements(measurements []transport.Measurement) {

    if time.Since(lastMessageDisplay).Seconds() < 5.0 {
        return
    }

    defer func() { lastMessageDisplay = time.Now() }()

    if lastMessageDisplay.IsZero() {
        sub.StatusMessage("Receiving measurements...")
        return
    }

    var message strings.Builder

	message.WriteString(format.UInt64(sub.TotalMeasurementsReceived()))
    message.WriteString(" measurements received so far...\n")
    message.WriteString("Timestamp: ")
    message.WriteString(measurements[0].DateTime().Format("2006-01-02 15:04:05.999999999"))
    message.WriteRune('\n')
    message.WriteString("\tID\tSignal ID\t\t\t\tValue\n")

    for i := 0; i < len(measurements); i++ {
        measurement := measurements[i]
        metadata := sub.Metadata(&measurement)

        message.WriteRune('\t')
        message.WriteString(strconv.FormatUint(metadata.ID, 10))
        message.WriteRune('\t')
        message.WriteString(measurement.SignalID.String())
        message.WriteRune('\t')
		message.WriteString(format.Float(measurement.Value, 6))
        message.WriteRune('\n')
    }

    sub.StatusMessage(message.String())
}

// ConnectionTerminated handles notification that a connection has been terminated.
func (sub *TestSubscriber) ConnectionTerminated() {
    // Call base implementation method which will display a connection terminated message to stderr
    sub.SubscriberBase.ConnectionTerminated()

    // Reset last message display time on disconnect
    lastMessageDisplay = time.Time{}
}
```

## More Examples
> [https://github.com/sttp/goapi/tree/main/examples](https://github.com/sttp/goapi/tree/main/examples)


## Quick Installation
```console
go get https://github.com/sttp/goapi
```

## Support
For discussion and support, join our [discussions channel](https://github.com/sttp/goapi/discussions) or [open an issue](https://github.com/sttp/goapi/issues) on GitHub.
## Links

* [STTP Go API Package Docs](https://pkg.go.dev/github.com/sttp/goapi)
* [STTP General Documentation](https://sttp.github.io/documentation/)
* [STTP (IEEE 2664) Standard](https://standards.ieee.org/project/2664.html)