workspace(name = "founding")

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "9b3a85e62cc8c00d4b356bfb2035ca0657703187",
)

load("@io_bazel_rules_go//go:def.bzl", "go_register_toolchains", "go_repository", "go_rules_dependencies")

go_rules_dependencies()
go_register_toolchains()

go_repository(
    name = "io_k8s_api",
    importpath = "k8s.io/api",
    commit = "5584376ceeffeb13a2e98b5e9f0e9dab37de4bab",
)

go_repository(
    name = "io_k8s_apimachinery",
    commit = "762f667215d26148452351d3d7de61dd8bb6d260",
    importpath = "k8s.io/apimachinery",
)

go_repository(
    name = "io_k8s_client_go",
    commit = "fff8c3d73ec5a59df69cd957ad8abeaf68f7c9be",
    importpath = "k8s.io/client-go",
)

