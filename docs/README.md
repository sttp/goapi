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
    defer subscriber.Close()

    // Start data read at each connection
    subscriber.SetConnectionEstablishedReceiver(func() {
        go func() {
            reader := subscriber.ReadMeasurements()
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
    })

    subscriber.Subscribe("FILTER TOP 20 ActiveMeasurements WHERE True", nil)
    subscriber.Dial("localhost:7175", nil)

    bufio.NewReader(os.Stdin).ReadRune()
}
```

Example Output:
```cmd
Connection to localhost:7175 established.
Received 28,323 bytes of metadata in 0.020 seconds. Decompressing...
Decompressed 251,898 bytes of metadata in 0.001 seconds. Parsing...
Parsed 532 metadata records in 0.023 seconds.
    Discovered:
        4 DeviceDetail records
        434 MeasurementDetail records
        93 PhasorDetail records
        1 SchemaVersion records
Metadata schema version: 14
Received success code in response to server command: Subscribe
Client subscribed as compact with 20 signals.
Receiving measurements...
2980 measurements received so far. Current measurement:
    {bcc6b18e-ed62-4c93-bc55-c7060ff58d5e} @ 16:08:14.366 = 0.036 (Normal)
5960 measurements received so far. Current measurement:
    {76cf5782-72f3-4312-ab92-d1e04bfd0e80} @ 16:08:19.366 = 155.645 (Normal)
8922 measurements received so far. Current measurement:
    {bcc6b18e-ed62-4c93-bc55-c7060ff58d5e} @ 16:08:24.4 = -3.473 (Normal)
11896 measurements received so far. Current measurement:
    {bcc6b18e-ed62-4c93-bc55-c7060ff58d5e} @ 16:08:29.366 = -0.434 (Normal)

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