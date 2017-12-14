workspace(name = "bazel_bugs")

http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.8.1/rules_go-0.8.1.tar.gz",
    sha256 = "90bb270d0a92ed5c83558b2797346917c46547f6f7103e648941ecdb6b9d0e72",
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
