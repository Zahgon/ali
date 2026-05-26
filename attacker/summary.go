package attacker

import (
	"time"

	"github.com/nakabonne/ali/export"
)

func newSummary(targetURL, method string, rate int, duration time.Duration, metrics *Metrics) export.Summary {
	_ = "STUB: not implemented"
	return *new(export.Summary)
}

func durationToMillis(d time.Duration) float64 { _ = "STUB: not implemented"; return 0 }
