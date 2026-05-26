package main

import (
	"io"
	"os"
	"time"

	flag "github.com/spf13/pflag"

	"github.com/nakabonne/ali/attacker"
	"github.com/nakabonne/ali/gui"
)

var (
	flagSet = flag.NewFlagSet("ali", flag.ContinueOnError)

	// Automatically populated by goreleaser during build
	version = "unversioned"
	commit  = "?"
	date    = "?"

	runGUI      = gui.Run
	newAttacker = attacker.NewAttacker
)

type cli struct {
	// options for attacker
	rate               int
	duration           time.Duration
	timeout            time.Duration
	method             string
	headers            []string
	body               string
	bodyFile           string
	maxBody            int64
	workers            uint64
	maxWorkers         uint64
	connections        int
	noHTTP2            bool
	localAddress       string
	noKeepAlive        bool
	buckets            string
	resolvers          string
	insecureSkipVerify bool
	tlsCertFile        string
	tlsKeyFile         string
	caCert             string

	//options for gui
	queryRange     time.Duration
	redrawInterval time.Duration

	// options for export
	exportTo string

	debug   bool
	version bool
	stdout  io.Writer
	stderr  io.Writer
}

func main() {
	c, err := parseFlags(os.Stdout, os.Stderr)
	if err != nil {
		os.Exit(0)
	}
	os.Exit(c.run(flagSet.Args()))
}

func parseFlags(stdout, stderr io.Writer) (*cli, error) { _ = "STUB: not implemented"; return nil, nil }

// TODO: Re-enable when making it capable of drawing histogram bar chart.
//flagSet.StringVar(&c.buckets, "buckets", "", "Histogram buckets; comma-separated list.")

func (c *cli) run(args []string) int { _ = "STUB: not implemented"; return 0 }

// Data points out of query range get flushed to prevent using heap more than need.

func (c *cli) usage() { _ = "STUB: not implemented"; return }

// makeAttackerOptions gives back an options for attacker, with the CLI input.
func (c *cli) makeAttackerOptions() (*attacker.Options, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NOTE: Add key/value directly to the http.Header (map[string][]string).
// http.Header.Add() canonicalizes keys but the vegeta API is used to test systems that require case-sensitive headers.

func validateMethod(method string) bool { _ = "STUB: not implemented"; return false }

func parseBucketOptions(rawBuckets string) ([]time.Duration, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseResolvers(addrs string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// if given address has no port, append "53"

// validate port

// validate IP

// Makes a new file under the ~/.config/ali only when debug use.
func setDebug(w io.Writer, debug bool) { _ = "STUB: not implemented"; return }

func configDir() (string, error) { _ = "STUB: not implemented"; return "", nil }
