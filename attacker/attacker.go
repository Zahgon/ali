package attacker

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"math"
	"net"
	"net/http"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"

	"github.com/nakabonne/ali/export"
	"github.com/nakabonne/ali/storage"
)

const (
	DefaultRate        = 50
	DefaultDuration    = 10 * time.Second
	DefaultTimeout     = 30 * time.Second
	DefaultMethod      = http.MethodGet
	DefaultWorkers     = 10
	DefaultMaxWorkers  = math.MaxUint64
	DefaultMaxBody     = int64(-1)
	DefaultConnections = 10000
)

var DefaultLocalAddr = net.IPAddr{IP: net.IPv4zero}

// Options provides optional settings to attack.
type Options struct {
	Rate        int
	Duration    time.Duration
	Timeout     time.Duration
	Method      string
	Body        []byte
	MaxBody     int64
	Header      http.Header
	Workers     uint64
	MaxWorkers  uint64
	KeepAlive   bool
	Connections int
	HTTP2       bool
	LocalAddr   net.IPAddr
	Buckets     []time.Duration
	Resolvers   []string

	InsecureSkipVerify bool
	CACertificatePool  *x509.CertPool
	TLSCertificates    []tls.Certificate

	Attacker backedAttacker

	Exporter    *export.FileExporter
	IDGenerator func() string
}

type Attacker interface {
	// Attack keeps the request running for the specified period of time.
	// Results are sent to the given channel as soon as they arrive.
	// When the attack is over, it gives back final statistics.
	// TODO: Use storage instead of metricsCh
	Attack(ctx context.Context, metricsCh chan *Metrics) error

	// Rate gives back the rate set to itself.
	Rate() int
	// Rate gives back the duration set to itself.
	Duration() time.Duration
	// Rate gives back the method set to itself.
	Method() string
}

func NewAttacker(storage storage.Writer, target string, opts *Options) (Attacker, error) {
	_ = "STUB: not implemented"
	return *new(Attacker), nil
}

type backedAttacker interface {
	Attack(vegeta.Targeter, vegeta.Pacer, time.Duration, string) <-chan *vegeta.Result
	Stop()
}

type attacker struct {
	target             string
	rate               int
	duration           time.Duration
	timeout            time.Duration
	method             string
	body               []byte
	maxBody            int64
	header             http.Header
	workers            uint64
	maxWorkers         uint64
	keepAlive          bool
	connections        int
	http2              bool
	localAddr          net.IPAddr
	buckets            []time.Duration
	resolvers          []string
	insecureSkipVerify bool
	caCertificatePool  *x509.CertPool
	tlsCertificates    []tls.Certificate

	attacker backedAttacker
	storage  storage.Writer

	exporter    *export.FileExporter
	idGenerator func() string
}

func (a *attacker) Attack(ctx context.Context, metricsCh chan *Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *attacker) Rate() int { _ = "STUB: not implemented"; return 0 }

func defaultIDGenerator() string { _ = "STUB: not implemented"; return "" }

func (a *attacker) Duration() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (a *attacker) Method() string { _ = "STUB: not implemented"; return "" }
