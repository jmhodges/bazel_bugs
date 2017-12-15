workspace(name = "bazel_bugs")

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "b2a59d8140f33174ca9cbac2cf5ab0bf0997826c",
)

load("@io_bazel_rules_go//go:def.bzl", "go_register_toolchains", "go_repository", "go_rules_dependencies")

go_rules_dependencies()
go_register_toolchains()

git_repository(
    name = "io_bazel_rules_webtesting",
    remote = "https://github.com/bazelbuild/rules_webtesting.git",
    commit = "d9009f881c19d3da520943955d9df80b492e9235",
)

load("@io_bazel_rules_webtesting//web:repositories.bzl",
    "browser_repositories",
    "web_test_repositories")


web_test_repositories()

browser_repositories(
    chromium = True,
)
