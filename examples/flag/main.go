package main

import (
	"flag"
	"fmt"

	"github.com/gebn/go-stamp/v2"
)

func main() {
	version := flag.Bool("version", false, "Print version and exit")
	flag.Parse()

	if *version {
		fmt.Println(stamp.Summary())
		return
	}
}
