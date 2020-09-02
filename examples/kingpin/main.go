package main

import (
	"github.com/alecthomas/kingpin"
	"github.com/gebn/go-stamp/v2"
)

func main() {
	kingpin.Version(stamp.Summary())
	kingpin.Parse()
}
