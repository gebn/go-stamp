# Stamp

[![CI](https://github.com/gebn/go-stamp/actions/workflows/build.yaml/badge.svg)](https://github.com/gebn/go-stamp/actions/workflows/build.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/gebn/go-stamp/v2.svg)](https://pkg.go.dev/github.com/gebn/go-stamp/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/gebn/go-stamp)](https://goreportcard.com/report/github.com/gebn/go-stamp)

Stamp exposes information about the build (user, host, time) and code version (description, commit, branch) at runtime.

## Configuration

To include the library, add the following to your `WORKSPACE`:

```python
load("@bazel_gazelle//:deps.bzl", "go_repository")
go_repository(
    name = "com_github_gebn_go_stamp_v2",
    tag = "v2.2.0",
    importpath = "github.com/gebn/go-stamp/v2",
)
```

Stamp uses a combination of [`workspace_status_command`](https://docs.bazel.build/versions/master/user-manual.html#flag--workspace_status_command) and [`x_defs`](https://github.com/bazelbuild/rules_go/blob/master/go/core.rst#id22) to substitute values of Go `var`s at link time. Unfortunately, there does not currently appear to be a way to add to the workspace status from Starlark, so this must be manually configured.

First, create an executable script with the following contents somewhere in your workspace, e.g. `bin/workspace_status` (if you already have a workspace status script, you just need to add the three `echo` lines to it). N.B. due to [this](https://github.com/bazelbuild/bazel/issues/5002) issue, having Bazel execute a script in the workspace root is awkward.

```bash
#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

echo "STABLE_STAMP_COMMIT $(git rev-parse HEAD)"
echo "STABLE_STAMP_BRANCH $(git rev-parse --abbrev-ref HEAD)"  # will be HEAD in detached HEAD state
echo "STABLE_STAMP_VERSION $(git describe --always --tags --dirty)"
```

Then add the following line to `.bazelrc` to execute the script during the build:

    build --workspace_status_command=bin/workspace_status --stamp

You can test everything is working by calling `Summary()` in your project, which summarises all information gathered by this library:

```go
package main

import (
    "fmt"

    "github.com/gebn/go-stamp/v2"
)

func main() {
    fmt.Println(stamp.Summary())
}
```

This should print something similar to:

    v1.0.0 (709d67c5f0563c685838312cb33a2a92ab1788f5, master), built with go1.13.1 by george@dev on 2019-10-05T18:09:58+01:00

See [`examples`](examples/) for more detail.
[`vanilla`](examples/vanilla/) demos using this library without Bazel.
