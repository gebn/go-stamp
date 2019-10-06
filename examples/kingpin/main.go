package main

import (
	"github.com/alecthomas/kingpin"
	"github.com/gebn/go-stamp"
)

func main() {
	kingpin.Version(stamp.Summary())
	kingpin.Parse()
}
