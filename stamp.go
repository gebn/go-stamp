// Package stamp provides build-time workspace status information at runtime,
// including the revision version and build metadata.
package stamp

import (
	"fmt"

	"github.com/gebn/go-stamp/build"
	"github.com/gebn/go-stamp/revision"
)

// String returns a human-readable summary of the revision version and build
// metadata, e.g. "1.0.0-1-g157ed0b (157ed0bb7b7de3c4c2e750a5b9ee675e2997ea80,
// master), built with go1.11.5 by george@dev on 2019-02-01T22:45:21Z".
func String() string {
	return fmt.Sprintf("%v, %v", revision.String(), build.String())
}
