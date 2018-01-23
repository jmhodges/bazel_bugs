workspace(name = "bazel_bugs")

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "b2a59d8140f33174ca9cbac2cf5ab0bf0997826c",
)

git_repository(
    name = "bazel_skylib",
    remote = "https://github.com/bazelbuild/bazel-skylib.git",
    commit = "ff23a62c57d2912c3073a69c12f42c3d6e58a957",
)

load("@io_bazel_rules_go//go:def.bzl", "go_register_toolchains", "go_repository", "go_rules_dependencies")

go_rules_dependencies()
go_register_toolchains()

git_repository(
    name = "io_bazel_rules_webtesting",
    remote = "https://github.com/bazelbuild/rules_webtesting.git",
    commit = "980c8ab6f4869d4315f7f155c55ecd5495740630",
)

load("@io_bazel_rules_webtesting//web:repositories.bzl",
    "browser_repositories",
    "web_test_repositories")


web_test_repositories()

browser_repositories(
    chromium = True,
)
