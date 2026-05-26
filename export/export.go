package export

import (
	"bufio"
	"encoding/csv"
	"os"
	"time"
)

const (
	resultsFilename = "results.csv"
)

var resultsHeader = []string{"id", "timestamp", "latency_ns", "url", "method", "status_code"}

type Meta struct {
	ID        string
	TargetURL string
	Method    string
	Rate      int
	Duration  time.Duration
}

type Result struct {
	Timestamp  time.Time
	LatencyNS  float64
	URL        string
	Method     string
	StatusCode uint16
}

type Summary struct {
	Target      TargetSummary      `json:"target"`
	Parameters  ParametersSummary  `json:"parameters"`
	Timing      TimingSummary      `json:"timing"`
	Requests    RequestsSummary    `json:"requests"`
	Throughput  float64            `json:"throughput"`
	LatencyMS   LatencySummary     `json:"latency_ms"`
	Bytes       BytesSummary       `json:"bytes"`
	StatusCodes StatusCodesSummary `json:"status_codes"`
}

type TargetSummary struct {
	URL    string `json:"url"`
	Method string `json:"method"`
}

type ParametersSummary struct {
	Rate            int     `json:"rate"`
	DurationSeconds float64 `json:"duration_seconds"`
}

type TimingSummary struct {
	Earliest time.Time `json:"earliest"`
	Latest   time.Time `json:"latest"`
}

type RequestsSummary struct {
	Count        uint64  `json:"count"`
	SuccessRatio float64 `json:"success_ratio"`
}

type LatencySummary struct {
	Total float64 `json:"total"`
	Mean  float64 `json:"mean"`
	P50   float64 `json:"p50"`
	P90   float64 `json:"p90"`
	P95   float64 `json:"p95"`
	P99   float64 `json:"p99"`
	Max   float64 `json:"max"`
	Min   float64 `json:"min"`
}

type BytesSummary struct {
	In  BytesFlowSummary `json:"in"`
	Out BytesFlowSummary `json:"out"`
}

type BytesFlowSummary struct {
	Total uint64  `json:"total"`
	Mean  float64 `json:"mean"`
}

type StatusCodesSummary map[string]int

func (s StatusCodesSummary) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type FileExporter struct {
	dir string
}

func NewFileExporter(dir string) *FileExporter { _ = "STUB: not implemented"; return nil }

type Run struct {
	meta Meta

	resultsPath string
	summaryPath string

	resultsFile *os.File
	resultsBuf  *bufio.Writer
	resultsCSV  *csv.Writer

	tempResultsPath string
	closed          bool
}

func (e *FileExporter) StartRun(meta Meta) (*Run, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Run) WriteResult(res Result) error { _ = "STUB: not implemented"; return nil }

func (r *Run) Close(summary Summary) error { _ = "STUB: not implemented"; return nil }

func (r *Run) Abort() error { _ = "STUB: not implemented"; return nil }

func writeSummary(path string, summary Summary) error { _ = "STUB: not implemented"; return nil }

func formatLatencyNS(v float64) string { _ = "STUB: not implemented"; return "" }

func summaryFilename(id string) string { _ = "STUB: not implemented"; return "" }
