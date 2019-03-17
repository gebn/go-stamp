package stamp

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

var (
	// User is the username of the account that built the library.
	User string

	// Host is the unqualified hostname of the machine the library was
	// built on.
	Host string

	// Timestamp is when the library was built, represented as the
	// number of seconds since the epoch. Use Time() to get this value
	// in a form that is easier to work with.
	timestamp string
)

// Time returns the build time string in its native representation. If
// the build time was not provided, returns the zero time instant.
func Time() time.Time {
	i, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return time.Time{}
	}
	return time.Unix(i, 0)
}

// BuildSummary returns a human-readable summary of the build metadata,
// e.g. "built with go1.11.5 by george@dev on 2019-02-01T22:45:21Z".
func BuildSummary() string {
	return fmt.Sprintf("built with %v by %v@%v on %v", runtime.Version(),
		User, Host, Time().Format(time.RFC3339))
}
