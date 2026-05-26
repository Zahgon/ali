package gui

import (
	"context"
	"sync"
	"time"

	"go.uber.org/atomic"

	"github.com/nakabonne/ali/attacker"
	"github.com/nakabonne/ali/storage"
)

// drawer periodically queries data points from the storage and passes them to the termdash API.
type drawer struct {
	// specify the data points range to show on the UI
	queryRange     time.Duration
	redrawInterval time.Duration
	widgets        *widgets
	gridOpts       *gridOpts

	metricsCh chan *attacker.Metrics

	// aims to avoid to perform multiple `appendChartValues`.
	chartDrawing *atomic.Bool

	mu      sync.RWMutex
	metrics *attacker.Metrics
	storage storage.Reader

	errMu     sync.Mutex
	exportErr error
}

// redrawCharts sets the values held by itself as chart values, at the specified interval as redrawInterval.
func (d *drawer) redrawCharts(ctx context.Context) { _ = "STUB: not implemented"; return }

func (d *drawer) redrawGauge(ctx context.Context, duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

// as time.Duration is the unit of nanoseconds
// small duration can exceed 100 on slow machines

const (
	latenciesTextFormat = `Total: %v
Mean: %v
P50: %v
P90: %v
P95: %v
P99: %v
Max: %v
Min: %v`

	bytesTextFormat = `In:
  Total: %v
  Mean: %v
Out:
  Total: %v
  Mean: %v`

	othersTextFormat = `Duration: %v
Wait: %v
Requests: %d
Rate: %f
Throughput: %f
Success: %f
Earliest: %v
Latest: %v
End: %v`
)

// redrawMetrics writes the metrics held by itself into the widgets, at the specified interval as redrawInterval.
func (d *drawer) redrawMetrics(ctx context.Context) { _ = "STUB: not implemented"; return }

// To guarantee that status codes are in order
// taking the slice of keys and sorting them.

func (d *drawer) updateMetrics(ctx context.Context) { _ = "STUB: not implemented"; return }

func (d *drawer) setExportErr(err error) { _ = "STUB: not implemented"; return }

func (d *drawer) exportError() error { _ = "STUB: not implemented"; return nil }
