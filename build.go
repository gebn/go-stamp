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

	// timestamp is when the library was built, represented as the
	// number of seconds since the epoch. Use Time() to get this value
	// in a form that is easier to work with.
	timestamp string
)

// Time returns the build time. Returns the zero time instant if the
// build time is unavailable.
func Time() time.Time {
	// the timestamp is really System.currentTimeMillis() / 1000, which
	// is a 64-bit type, so we parse it as such rather than with Atoi()
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
