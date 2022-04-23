# Flag

This example shows how to use go-stamp with [`flag`](https://pkg.go.dev/flag) such that passing `--version` prints the version summary.

## Usage

```
$ bazel run //examples/flag -- --version
INFO: Analyzed target //examples/flag:flag (0 packages loaded, 0 targets configured).
INFO: Found 1 target...
Target //examples/flag:flag up-to-date:
  bazel-bin/examples/flag/linux_amd64_stripped/flag
INFO: Elapsed time: 0.208s, Critical Path: 0.03s
INFO: 0 processes.
INFO: Build completed successfully, 1 total action
INFO: Build completed successfully, 1 total action
v2.0.2 (d37c3b01740c3e9f9ee86fef86a52e503f103ec5, master), built with go1.13.1 by george@dev on 2019-10-06T10:24:30+01:00
```
