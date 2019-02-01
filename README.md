# Stamp

[![GoDoc](https://godoc.org/github.com/gebn/go-stamp?status.svg)](https://godoc.org/github.com/gebn/go-stamp)
[![Go Report Card](https://goreportcard.com/badge/github.com/gebn/go-stamp)](https://goreportcard.com/report/github.com/gebn/go-stamp)

Stamp exposes information about the build (user, fqdn, time) and code version (description, commit, branch) at runtime.

## Configuration

To include the library, add the following to your `WORKSPACE`:

    load("@bazel_gazelle//:deps.bzl", "go_repository")
    go_repository(
        name = "com_github_gebn_go_stamp",
        commit = "810b8c5846be228396fcc85dba20908c49ef4d77",
        importpath = "github.com/gebn/go-stamp",
    )

Stamp uses [`workspace_status_command`](https://docs.bazel.build/versions/master/user-manual.html#flag--workspace_status_command) to set variables to the correct value at link time.
Unfortunately there does not currently appear to be a way to add keys to the workspace status within a library, so this must be set up on the root project.
Create an executable script with the following contents somewhere in your project, e.g. `bin/workspace_status` (if you already have this set up, you just need to add the three `echo` lines to your existing script):

    #!/bin/bash

    set -o errexit
    set -o nounset
    set -o pipefail

    echo "STABLE_GIT_DESCRIBE $(git describe --always --tags --dirty)"
    echo "STABLE_GIT_COMMIT $(git rev-parse HEAD)"
    echo "STABLE_GIT_BRANCH $(git rev-parse --abbrev-ref HEAD)"  # will be HEAD in detached HEAD state

Then add the following line to `.bazelrc` to execute the script during the build:

    build --workspace_status_command=bin/workspace_status

You can test everything is working by calling `String()` in your project, which summarises all information gathered by this library:

    package main

    import (
        "fmt"
        "github.com/gebn/go-stamp"
    )

    func main() {
        fmt.Println(stamp.String())
    }

This should print something similar to:

    v1.0.0-11-gcb25cf1 (cb25cf163fdb9f4d5033b9a4d70e9b259ef4b9b2, master), built with go1.11.5 by george@dev on 2019-02-01T21:07:53Z

### Kingpin

If using [Kingpin](https://github.com/alecthomas/kingpin), Stamp can be integrated like this:

    func main() {
        kingpin.Version(stamp.String())
        kingpin.Parse()
        // ...
    }
