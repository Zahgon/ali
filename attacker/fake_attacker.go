package attacker

import (
	"context"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

type FakeAttacker struct {
	rate     int
	duration time.Duration
	method   string
}

func (f *FakeAttacker) Attack(ctx context.Context, metricsCh chan *Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *FakeAttacker) Rate() int { _ = "STUB: not implemented"; return 0 }

func (f *FakeAttacker) Duration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (f *FakeAttacker) Method() string { _ = "STUB: not implemented"; return "" }

type fakeBackedAttacker struct {
	results []*vegeta.Result
}

func (f *fakeBackedAttacker) Attack(vegeta.Targeter, vegeta.Pacer, time.Duration, string) <-chan *vegeta.Result {
	_ = "STUB: not implemented"
	return nil
}

func (f *fakeBackedAttacker) Stop() { _ = "STUB: not implemented"; return }
