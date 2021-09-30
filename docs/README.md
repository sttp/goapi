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
    "os"
    "strconv"
    "strings"
    "time"

    "github.com/sttp/goapi/sttp"
)

func main() {
    subscriber := sttp.NewSubscriber()
    subscriber.Subscribe("FILTER TOP 20 ActiveMeasurements WHERE True", nil)
    subscriber.Dial("localhost:7175", nil)
    defer subscriber.Close()

    reader := subscriber.ReadMeasurements()

    go func() {
        var lastMessage time.Time

        for subscriber.IsConnected() {
            measurement := reader.NextMeasurement()

            if time.Since(lastMessage).Seconds() < 5.0 {
                continue
            } else if lastMessage.IsZero() {
                subscriber.StatusMessage("Receiving measurements...")
                lastMessage = time.Now()
                continue
            }

            var message strings.Builder

            message.WriteString(strconv.FormatUint(subscriber.TotalMeasurementsReceived(), 10))
            message.WriteString(" measurements received so far. Current measurement:\n    ")
            message.WriteString(measurement.String())

            subscriber.StatusMessage(message.String())
            lastMessage = time.Now()
        }
    }()

    bufio.NewReader(os.Stdin).ReadRune()
}
```

Example Output:
```cmd
Connection to localhost:7175 established.
Received 34,884 bytes of metadata in 0.586 seconds. Decompressing...
Decompressed 304,329 bytes of metadata in 0.003 seconds. Parsing...
Parsed 643 metadata records in 1.236 seconds.
    Discovered:
        4 DeviceDetail records
        517 MeasurementDetail records
        121 PhasorDetail records
        1 SchemaVersion records
Metadata schema version: 14
Received success code in response to server command 0x2
Client subscribed as compact with 20 signals.
Receiving measurements...
2960 measurements received so far. Current measurement:
    {7c30766c-8775-4763-b4e4-b2bb9512f464} @ 07:55:29.333 = 0.000 (Norm)
5900 measurements received so far. Current measurement:
    {f395d119-9b56-433b-9368-c78221d22639} @ 07:55:34.333 = 7.000 (Norm)
8860 measurements received so far. Current measurement:
    {f395d119-9b56-433b-9368-c78221d22639} @ 07:55:39.333 = 7.000 (Norm)
11840 measurements received so far. Current measurement:
    {f395d119-9b56-433b-9368-c78221d22639} @ 07:55:44.366 = 7.000 (Norm)
14814 measurements received so far. Current measurement:
    {3fccba58-ce36-4b05-8418-4cd9152aa9f0} @ 07:55:49.366 = 10.000 (Norm)

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