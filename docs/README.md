## Go STTP ([IEEE 2664](https://standards.ieee.org/project/2664.html)) Documentation
### Streaming Telemetry Transport Protocol

<img align="right" src="img/sttp.png">

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

Example Output:
```
Connection to localhost:7175 established.
Received 28,323 bytes of metadata in 0.017 seconds. Decompressing...
Decompressed 251,898 bytes of metadata in 0.001 seconds. Parsing...
Parsed 532 metadata records in 0.021 seconds.
    Discovered:
        4 DeviceDetail records
        434 MeasurementDetail records
        93 PhasorDetail records
        1 SchemaVersion records
Metadata schema version: 14
Received signal index cache with 20 mappings
Received success code in response to server command 0x2
Client subscribed as compact with 20 signals.
Receiving measurements...
2,994 measurements received so far...
Timestamp: 2021-09-29 13:19:52.4333333
        ID      Signal ID                               Value
        152     {76cf5782-72f3-4312-ab92-d1e04bfd0e80}  149.011917
        155     {bcc6b18e-ed62-4c93-bc55-c7060ff58d5e}  -2.539322
        153     {fefec5df-ca04-4c2b-a7b2-b4c8b1298795}  42.401386
        156     {2c9a565f-424c-44c6-ac03-c6f8be199b24}  48.603382
        154     {f0a6f8c5-0c0b-48d4-b181-db45ed555b7e}  -67.139977
        157     {6ee8c6ca-3421-4867-846e-b301f730702e}  -10.387241

5,968 measurements received so far...
Timestamp: 2021-09-29 13:19:57.4333333
        ID      Signal ID                               Value
        152     {76cf5782-72f3-4312-ab92-d1e04bfd0e80}  150.768799
        155     {bcc6b18e-ed62-4c93-bc55-c7060ff58d5e}  -1.029660
        153     {fefec5df-ca04-4c2b-a7b2-b4c8b1298795}  42.121387
        156     {2c9a565f-424c-44c6-ac03-c6f8be199b24}  48.389652
        154     {f0a6f8c5-0c0b-48d4-b181-db45ed555b7e}  -67.775978
        157     {6ee8c6ca-3421-4867-846e-b301f730702e}  -11.303692

Connection to localhost:7175 terminated.
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