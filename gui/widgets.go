package gui

import (
	"time"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgetapi"
	"github.com/mum4k/termdash/widgets/gauge"
	"github.com/mum4k/termdash/widgets/linechart"
	"github.com/mum4k/termdash/widgets/text"
)

type LineChart interface {
	widgetapi.Widget
	Series(label string, values []float64, opts ...linechart.SeriesOption) error
}

type Text interface {
	widgetapi.Widget
	Write(text string, wOpts ...text.WriteOption) error
}

type Gauge interface {
	widgetapi.Widget
	Percent(p int, opts ...gauge.Option) error
}

type chartLegend struct {
	text     Text
	cellOpts []cell.Option
}

type widgets struct {
	latencyChart LineChart

	paramsText      Text
	latenciesText   Text
	bytesText       Text
	statusCodesText Text
	errorsText      Text
	othersText      Text

	percentilesChart LineChart
	p99Legend        chartLegend
	p95Legend        chartLegend
	p90Legend        chartLegend
	p50Legend        chartLegend

	progressGauge Gauge
	navi          Text
}

// Thg given params is used for displayed text.
func newWidgets(targetURL string, rate int, duration time.Duration, method string) (*widgets, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newLineChart() (LineChart, error) { _ = "STUB: not implemented"; return *new(LineChart), nil }

func newText(s string, opts ...text.WriteOption) (Text, error) {
	_ = "STUB: not implemented"
	return *new(Text), nil
}

func newGauge() (Gauge, error) { _ = "STUB: not implemented"; return *new(Gauge), nil }

func makeParamsText(targetURL string, rate int, duration time.Duration, method string) string {
	_ = "STUB: not implemented"
	return ""
}
