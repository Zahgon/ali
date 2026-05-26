package storage

import (
	"time"

	"github.com/nakabonne/tstorage"
)

const (
	LatencyMetricName = "latency"
	P50MetricName     = "p50"
	P90MetricName     = "p90"
	P95MetricName     = "p95"
	P99MetricName     = "p99"
)

// Storage provides goroutine safe capabilities of insertion into and retrieval from the time-series storage.
// Backed by "nakabonne/tstorage"
type Storage interface {
	Writer
	Reader
}

type Writer interface {
	Insert(result *Result) error
}

type Reader interface {
	Select(metric string, start, end time.Time) ([]float64, error)
}

// Result contains the results of a single HTTP request.
type Result struct {
	Code      uint16
	Timestamp time.Time
	Latency   time.Duration
	P50       time.Duration
	P90       time.Duration
	P95       time.Duration
	P99       time.Duration
}

func NewStorage(partitionDuration time.Duration) (Storage, error) {
	_ = "STUB: not implemented"
	return *new(Storage), nil
}

type storage struct {
	backend tstorage.Storage
}

// Insert writes the given result to the backend storage.
// The unit of value will be converted in milliseconds.
func (s *storage) Insert(result *Result) error {
	_ = "STUB: not implemented"
	// Convert timestamp into unix time in nanoseconds.
	return nil
}

// TODO: Think about how to handle code
/*
	labels := []tstorage.Label{
		{
			Name:  codeLabelName,
			Value: strconv.Itoa(int(result.Code)),
		},
	}
*/

func (s *storage) Select(metric string, start, end time.Time) ([]float64, error) {
	_ = "STUB: not implemented"
	// Convert timestamp into unix time in nanoseconds.
	return nil, nil
}
