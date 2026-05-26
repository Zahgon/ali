package gui

import (
	"context"

	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/terminal/terminalapi"

	"github.com/nakabonne/ali/attacker"
)

func navigateCharts(chartFuncs []func()) func(bool) { _ = "STUB: not implemented"; return nil }

func keybinds(ctx context.Context, cancel context.CancelFunc, c *container.Container, dr *drawer, a attacker.Attacker) func(*terminalapi.Keyboard) {
	_ = "STUB: not implemented"
	return nil
}

// Quit

// Attack

// backwards

// forwards

func attack(ctx context.Context, cancelParent context.CancelFunc, d *drawer, a attacker.Attacker) {
	_ = "STUB: not implemented"
	return
}

// To initialize, run redrawChart on a per-attack basis.
