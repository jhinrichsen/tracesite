package main

import (
	"fmt"
	"os"

	"github.com/jhinrichsen/tracesite/cli"
)

func main() {
	if err := cli.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
