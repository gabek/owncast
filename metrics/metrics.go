package metrics

import (
	"time"
)

// How often we poll for updates
const metricsPollingInterval = 15 * time.Second

type metrics struct {
	CPUUtilizations  []timestampedValue
	RAMUtilizations  []timestampedValue
	DiskUtilizations []timestampedValue

	Viewers []timestampedValue
}

// Metrics is the shared Metrics instance
var Metrics *metrics

// Start will begin the metrics collection and alerting
func Start() {
	Metrics = new(metrics)
	go startViewerCollectionMetrics()

	for range time.Tick(metricsPollingInterval) {
		handlePolling()
	}
}

func handlePolling() {
	// Collect hardware stats
	collectCPUUtilization()
	collectRAMUtilization()
	collectDiskUtilization()

	// Alerting
	handleAlerting()
}
