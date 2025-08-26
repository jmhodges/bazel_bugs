workspace(name = "com_github_jmhodges_bazel_bugs")

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_register_toolchains(
    version = "1.25.0",
)

go_rules_dependencies()

gazelle_dependencies()

go_repository(
    name = "com_google_cloud_go_bigquery",
    importpath = "cloud.google.com/go/bigquery",
    sum = "h1:rZvHnjSUs5sHK3F9awiuFk2PeOaB8suqNuim21GbaTc=",
    version = "v1.69.0",
)
