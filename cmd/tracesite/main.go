package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jhinrichsen/tracesite"
)

// Default values for all available options.
const (
	DefaultPort       = 33434
	DefaultPacketSize = 52
	DefaultTimeoutMs  = 4000
	DefaultRetries    = 3
	DefaultDest       = "kalbhor.xyz"
	DefaultMaxHops    = 24
	DefaultStartTTL   = 1
)

func main() {
	flag.Usage = func() {
		// mimic urfave/cli output
		out := flag.CommandLine.Output()
		me := os.Args[0]
		msg := strings.Join([]string{
			"NAME:",
			fmt.Sprintf("  %s - trace the route to a site", me),
			"",
			"USAGE:",
			fmt.Sprintf("  Trace a site: %s --hop=3 "+
				"--timeout=2000 kalbhor.xyz", me),
			"",
			"GLOBAL OPTIONS:",
			"",
		}, "\n")
		fmt.Fprintf(out, msg)
		flag.PrintDefaults()
	}
	var o tracesite.Options
	flag.IntVar(&o.Hop, "hop", DefaultStartTTL, "start from a custom hop number")
	flag.IntVar(&o.MaxHops, "maxhops", DefaultMaxHops, "custom max hops")
	flag.IntVar(&o.Port, "port", DefaultPort, "custom port number")
	flag.IntVar(&o.Timeout, "timeout", DefaultTimeoutMs, "custom timeout in ms")
	flag.IntVar(&o.Retries, "retries", DefaultRetries, "custom retries")
	flag.IntVar(&o.PacketSize, "packetsize", DefaultPacketSize, "custom packet size")

	flag.Parse()
	for _, destination := range flag.Args() {
		if err := tracesite.Tracesite(destination, o); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
	}
}
