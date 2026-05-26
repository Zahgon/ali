package attacker

import (
	"net"
)

type resolver struct {
	idx   uint64
	addrs []string
}

func NewResolver(addrs []string) *net.Resolver { _ = "STUB: not implemented"; return nil }

func (r *resolver) address() string { _ = "STUB: not implemented"; return "" }
