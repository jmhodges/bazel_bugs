workspace(name = "com_github_jmhodges_bazel_bugs")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

rules_docker_version = "0d355f103272e33e7da40506e9d8307fe7bfcb26"
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "65fbde25d4edcfc2bc61c8870e112ac4dc04f9c20274fe2e3d2d6b7bad11223c",
    strip_prefix = "rules_docker-{v}".format(v = rules_docker_version),
    urls = ["https://github.com/bazelbuild/rules_docker/archive/{v}.tar.gz".format(v = rules_docker_version)],
)

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)
container_repositories()

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
)

container_pull(
    name = "mysql",
    registry = "index.docker.io",
    repository = "library/mysql",
    # digest for tag 5.7.25 on 2019-02-28
    digest = "sha256:8c15b2612051244d0a2b6ceb6f9bf82ddc0e909555c1067c098e5f935e2751a7",
)
