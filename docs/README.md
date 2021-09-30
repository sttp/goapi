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

    bufio.NewReader(os.Stdin).ReadRune()
}
```

Example Output:
```cmd
Connection to localhost:7175 established.
Received 28,323 bytes of metadata in 0.020 seconds. Decompressing...
Decompressed 251,898 bytes of metadata in 0.001 seconds. Parsing...
Parsed 532 metadata records in 0.038 seconds.
    Discovered:
        4 DeviceDetail records
        434 MeasurementDetail records
        93 PhasorDetail records
        1 SchemaVersion records
Metadata schema version: 14
Received success code in response to server command 0x2
Client subscribed as compact with 20 signals.
Receiving measurements...
2986 measurements received so far. Current measurement:
    {c84fcf2f-d83a-4cd8-9dc8-7c5292c4a068} @ 18:21:42.766 = 60.019 (Norm)
5924 measurements received so far. Current measurement:
    {bcc6b18e-ed62-4c93-bc55-c7060ff58d5e} @ 18:21:42.766 = 1.397 (Norm)
8892 measurements received so far. Current measurement:
    {bcc6b18e-ed62-4c93-bc55-c7060ff58d5e} @ 18:21:47.766 = -1.739 (Norm)
11880 measurements received so far. Current measurement:
    {bcc6b18e-ed62-4c93-bc55-c7060ff58d5e} @ 18:21:52.766 = -4.064 (Norm)
14866 measurements received so far. Current measurement:
    {bcc6b18e-ed62-4c93-bc55-c7060ff58d5e} @ 18:21:57.766 = -2.472 (Norm)
17846 measurements received so far. Current measurement:
    {bcc6b18e-ed62-4c93-bc55-c7060ff58d5e} @ 18:22:02.766 = 0.569 (Norm)

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