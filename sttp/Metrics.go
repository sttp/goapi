package sttp

import "github.com/prometheus/client_golang/prometheus"

var (
	pmMetadataRefreshes     prometheus.Counter
	pmMetadataRefreshErrors prometheus.Counter

	pmMetadataRefreshPayloadSizes prometheus.Histogram
	pmMetadataRefreshDurations    prometheus.Histogram
)

func init() {
	pmMetadataRefreshes = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "sttp",
		Subsystem: "goapi",
		Name:      "metadata_refresh_cnt",
		Help:      "The number of metadata refreshes since program start",
	})

	pmMetadataRefreshErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "sttp",
		Subsystem: "goapi",
		Name:      "metadata_refresh_error_cnt",
		Help:      "The number of unsuccessful metadata refreshes since program start",
	})

	pmMetadataRefreshPayloadSizes = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "sttp",
		Subsystem: "goapi",
		Name:      "metadata_refresh_payload_sizes_bytes",
		Help:      "The sizes of observed metadata payloads in bytes",
		Buckets:   prometheus.ExponentialBuckets(float64(2>>14), 4.0, 8), // 16kb, 64kb, 256kb, 1MB, 4MB, 16MB, 64MB, 256MB, godzilla
	})

	pmMetadataRefreshDurations = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "sttp",
		Subsystem: "goapi",
		Name:      "metadata_refresh_duration_ms",
		Help:      "The duration of metadata refreshes in milliseconds",
	})

	prometheus.MustRegister(pmMetadataRefreshDurations, pmMetadataRefreshErrors, pmMetadataRefreshPayloadSizes, pmMetadataRefreshes)
}
