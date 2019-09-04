workspace(name = "com_github_jmhodges_bazel_bugs")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "b32e1087f5d66a6fcd81bbc389bc699947fdf81706ffaec955b8f4bdb629e4e7",
    strip_prefix = "rules_docker-165cd95365684f5fa1f29abc71cf1006e06c581a",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/165cd95365684f5fa1f29abc71cf1006e06c581a.zip"],

    # For 0.9.0:
    # sha256 = "e513c0ac6534810eb7a14bf025a0f159726753f97f74ab7863c650d26e01d677",
    # strip_prefix = "rules_docker-0.9.0",
    # urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.9.0/rules_docker-v0.9.0.tar.gz"],
)


load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

load("@io_bazel_rules_docker//container:new_pull.bzl", "new_container_pull")

new_container_pull(
    name = "mysql",
    registry = "index.docker.io",
    repository = "library/mysql",
    # digest for tag 5.7.25 on 2019-02-28
    digest = "sha256:8c15b2612051244d0a2b6ceb6f9bf82ddc0e909555c1067c098e5f935e2751a7",
)
