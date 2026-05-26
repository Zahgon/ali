package gui

import (
	"context"
	"time"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/terminal/terminalapi"

	"github.com/nakabonne/ali/attacker"
	"github.com/nakabonne/ali/storage"
)

const (
	DefaultQueryRange     = 30 * time.Second
	DefaultRedrawInterval = 250 * time.Millisecond
	minRedrawInterval     = 100 * time.Millisecond
	rootID                = "root"
	chartID               = "chart"
)

type Options struct {
	RedrawInternal time.Duration
	QueryRange     time.Duration
}

type runner func(ctx context.Context, t terminalapi.Terminal, c *container.Container, opts ...termdash.Option) error

func Run(targetURL string, storage storage.Reader, attacker attacker.Attacker, opts Options) error {
	_ = "STUB: not implemented"
	return nil
}

func run(t terminalapi.Terminal, r runner, targetURL string, storage storage.Reader, a attacker.Attacker, opts Options) error {
	_ = "STUB: not implemented"
	return nil
}

// newChartWithLegends creates a chart with legends at the bottom.
// TODO: use it for more charts than percentiles. Any chart that has multiple series would be able to use this func.
func newChartWithLegends(lineChart LineChart, opts []container.Option, texts ...Text) ([]container.Option, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// gridOpts holds all options in our grid.
// It basically holds the container options (column, width, padding, etc) of our widgets.
type gridOpts struct {
	// base options
	base []container.Option

	// so we can replace containers
	latency     []container.Option
	percentiles []container.Option
}

func gridLayout(w *widgets) (*gridOpts, error) { _ = "STUB: not implemented"; return nil, nil }
