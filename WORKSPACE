workspace(name = "com_github_jmhodges_bazel_bugs")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "49b04b1c7fe8636ca44455d72e2a5854bf666f21e50601383066c182d887ff58",
    strip_prefix = "rules_docker-4109d47b8656464e3075f3af1f6c8705e0eec867",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/4109d47b8656464e3075f3af1f6c8705e0eec867.zip"],

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

# load(
#     "@io_bazel_rules_docker//container:container.bzl",
#     "container_pull",
# )

load("@io_bazel_rules_docker//container:new_pull.bzl", "new_container_pull")

new_container_pull(
    name = "mysql",
    registry = "index.docker.io",
    repository = "library/mysql",
    # digest for tag 5.7.25 on 2019-02-28
    digest = "sha256:2bd4665d9c5ecad61f7ceff82f82e6470c4686b9ec0fd986b84012861506c722",
)
