package main

import (
	"fmt"
	"os"

	"github.com/jhinrichsen/tracesite/src/cli"
)

func main() {
	if err := cli.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
