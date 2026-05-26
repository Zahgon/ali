package storage

import "time"

type FakeStorage struct {
	Values []float64
	err    error
}

func (f *FakeStorage) Insert(_ *Result) error { _ = "STUB: not implemented"; return nil }

func (f *FakeStorage) Select(_ string, _, _ time.Time) ([]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
