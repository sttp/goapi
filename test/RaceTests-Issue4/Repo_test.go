package repro

import (
	"context"
	"log"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/sttp/goapi/sttp"
)

const (
	query  = "FILTER TOP 20 ActiveMeasurements WHERE SignalType <> 'STAT'"
	target = "127.0.0.1:7165" // 7165 default for openPDC, 7175 default for openHistorian
)

func run(dst string, query string, count int, instance int) {
	sub := sttp.NewSubscriber()
	conf := sttp.NewConfig()

	if instance < 0 {
		// Disable screen feedback for benchmarking
		sub.SetStatusMessageLogger(nil)
		sub.SetErrorMessageLogger(nil)
	} else {
		sub.SetStatusMessageLogger(func(msg string) {
			log.Println("Instance", instance, msg)
		})

		sub.SetErrorMessageLogger(func(msg string) {
			log.Println("Instance", instance, "error:", msg)
		})
	}

	sub.Subscribe(query, nil)

	conf.AutoRequestMetadata = false
	conf.MaxRetries = 3

	if err := sub.Dial(dst, conf); err != nil {
		log.Println("Instance", instance, err, "-- canceling instance")
		return
	}

	reader := sub.ReadMeasurements()
	const timeout = 5 // seconds
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	for sub.IsConnected() {
		if _, timeoutExpired := reader.NextMeasurement(ctx); timeoutExpired {
			log.Println("Instance", instance, "measurement read timed out after", timeout, "seconds -- canceling instance")
			break
		}

		count--

		if count <= 0 {
			break
		}

		if count%200 == 0 && instance > -1 {
			log.Println("Instance", instance, "received", count, "measurements so far...")
		}
	}

	sub.Close()
}

func TestRepro(t *testing.T) {
	const (
		instances    = 100
		measurements = 1000
	)

	wg := sync.WaitGroup{}
	wg.Add(instances)
	var count int32 = instances

	for i := 0; i < instances; i++ {
		go func(instance int) {
			run(target, query, measurements, instance)
			wg.Done()
			atomic.AddInt32(&count, -1)
			log.Println("Instance", instance, "done,", atomic.LoadInt32(&count), "remaining")
		}(i)
	}

	wg.Wait()
}

// Execute this bench test only: go test -bench ProcessMeasurements -run=^$ -count 5
func BenchmarkProcessMeasurements(b *testing.B) {
	run(target, query, 1000, -1)
}
