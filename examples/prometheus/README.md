# Prometheus

This example shows how to expose build information to Prometheus.
This is especially useful in aggregate to identify usage levels and old versions.

## Usage

Run the program with `bazel run //examples/prometheus`, then visit `localhost:8080`.
At the top, you should see two metrics resembling the following:

```
# HELP example_build_info The version and commit of the example code. Constant 1.
# TYPE example_build_info gauge
example_build_info{commit="d37c3b01740c3e9f9ee86fef86a52e503f103ec5",version="v2.0.1-1-gd37c3b0-dirty"} 1
# HELP example_build_time When the example code was build, as seconds since the Unix Epoch.
# TYPE example_build_time gauge
example_build_time 1.570322854e+09
```

For a less contrived example, see [`bmc_exporter`](https://github.com/gebn/bmc_exporter).
