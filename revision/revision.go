// Package revision embeds information about the revision contained
// in this library, including the commit identifier, branch name, and
// description of the version (e.g. tag name).
package revision

import (
	"fmt"
)

var (
	// Commit is the SHA-1 hash of the revision that was built to
	// produce this library.
	Commit string

	// Branch is the name of the branch the above commit was on at
	// the time the library was built. The CI pipeline checks out a
	// specific commit, meaning Git is in a detached HEAD state, so
	// this will always be "HEAD" for official releases.
	Branch string

	// Describe is the output of `git describe --always --tags --dirty`.
	// For official releases, this will be the tag name, which follows
	// semver, e.g. `v1.0.0`.
	Describe string
)

// String returns a human-readable summary of the revision metadata, e.g.
// "1.0.0-1-g157ed0b (157ed0bb7b7de3c4c2e750a5b9ee675e2997ea80, master)".
func String() string {
	return fmt.Sprintf("%v (%v, %v)", Describe, Commit, Branch)
}
