workspace(name = "com_github_jmhodges_bazel_bugs")

# required by gazelle
skylib_version = "8cecf885c8bf4c51e82fd6b50b9dd68d2c98f757"
http_archive(
    name = "bazel_skylib",
    sha256 = "68ef2998919a92c2c9553f7a6b00a1d0615b57720a13239c0e51d0ded5aa452a",
    strip_prefix= "bazel-skylib-{v}".format(v=skylib_version),
    urls = [
        "https://github.com/bazelbuild/bazel-skylib/archive/{v}.tar.gz".format(v=skylib_version)
    ]
)

load("@bazel_skylib//:lib.bzl", "versions")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.16.1/rules_go-0.16.1.tar.gz"],
    sha256 = "f5127a8f911468cd0b2d7a141f17253db81177523e4429796e14d429f5444f5f",
)

git_repository(
    name = "bazel_gazelle",
    commit = "f77eb9a55ac246a27f7a475c178740713bb45102",
    remote = "https://github.com/bazelbuild/bazel-gazelle.git"
)

load("@io_bazel_rules_go//go:def.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()
go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@bazel_gazelle//:def.bzl", "go_repository")
gazelle_dependencies()

git_repository(
    name = "io_bazel_rules_webtesting",
    remote = "https://github.com/bazelbuild/rules_webtesting.git",
    commit = "596d07c1f38486486969302158b9019418a5409e",
)

load("@io_bazel_rules_webtesting//web:repositories.bzl", "browser_repositories", "web_test_repositories")

browser_repositories(chromium = True)
web_test_repositories(omit_bazel_skylib = True)

