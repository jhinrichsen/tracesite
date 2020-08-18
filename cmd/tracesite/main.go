package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/jhinrichsen/tracesite"
)

const (
	DEFAULT_PORT        = 33434
	DEFAULT_PACKET_SIZE = 52
	DEFAULT_TIMEOUT_MS  = 4000
	DEFAULT_RETRIES     = 3
	DEFAULT_DEST        = "kalbhor.xyz"
	DEFAULT_MAX_HOPS    = 24
	DEFAULT_START_TTL   = 1
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
	flag.IntVar(&o.Hop, "hop", DEFAULT_START_TTL, "start from a custom hop number")
	flag.IntVar(&o.MaxHops, "maxhops", DEFAULT_MAX_HOPS, "custom max hops")
	flag.IntVar(&o.Port, "port", DEFAULT_PORT, "custom port number")
	flag.IntVar(&o.Timeout, "timeout", DEFAULT_TIMEOUT_MS, "custom timeout in ms")
	flag.IntVar(&o.Retries, "retries", DEFAULT_RETRIES, "custom retries")
	flag.IntVar(&o.PacketSize, "packetsize", DEFAULT_PACKET_SIZE, "custom packet size")

	flag.Parse()
	for _, destination := range flag.Args() {
		tracesite.Tracesite(destination, o)
	}
}
