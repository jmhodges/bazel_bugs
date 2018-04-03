workspace(name = "bazel_bugs")

http_archive(
    name = "bazel_skylib",
    sha256 = "587d706c4b19cf6ecb0516c86a674a7aa8f586a79c67fd56b4580372a7bac2ab",
    strip_prefix = "bazel-skylib-ff23a62c57d2912c3073a69c12f42c3d6e58a957",
    urls = [
        "https://github.com/bazelbuild/bazel-skylib/archive/ff23a62c57d2912c3073a69c12f42c3d6e58a957.tar.gz",
    ],
)

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "c003e0ccbe743017c87025917d6b8c42b1cc1442",
)

http_archive(
    name = "bazel_gazelle",
    url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.10.1/bazel-gazelle-0.10.1.tar.gz",
    sha256 = "d03625db67e9fb0905bbd206fa97e32ae9da894fe234a493e7517fd25faec914",
)

load("@io_bazel_rules_go//go:def.bzl", "go_register_toolchains", "go_repository", "go_rules_dependencies")

go_rules_dependencies()
go_register_toolchains()


load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()
