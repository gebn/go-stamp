load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.1/rules_go-0.18.1.tar.gz"],
    sha256 = "77dfd303492f2634de7a660445ee2d3de2960cbd52f97d8c0dffa9362d3ddef9",
)
http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
)
load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()

go_repository(
	name = "com_github_alecthomas_units",
	commit = "680d30ca31172657fa50e996eb82d790d1d8b96e",
	importpath = "github.com/alecthomas/units",
)

go_repository(
	name = "com_github_alecthomas_template",
	commit = "fb15b899a75114aa79cc930e33c46b577cc664b1",
	importpath = "github.com/alecthomas/template",
)

go_repository(
	name = "com_github_alecthomas_kingpin",
	importpath = "github.com/alecthomas/kingpin",
	tag = "v2.2.6",
)

go_repository(
	name = "com_github_prometheus_client_golang",
	importpath = "github.com/prometheus/client_golang",
	tag = "v1.1.0",
)

go_repository(
	name = "com_github_prometheus_common",
	importpath = "github.com/prometheus/common",
	tag = "v0.7.0",
)
