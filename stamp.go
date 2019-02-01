// Package stamp provides build-time workspace status information at runtime,
// including the revision version and build metadata.
package stamp

import (
	"fmt"

	"github.com/gebn/go-stamp/build"
	"github.com/gebn/go-stamp/revision"
)

// String returns a human-readable summary of the revision version and build
// metadata.
func String() string {
	return fmt.Sprintf("%v, %v", revision.String(), build.String())
}
