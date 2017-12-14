To test [bazel bug 2294](https://github.com/bazelbuild/bazel/issues/2294):

1. run `bazel build //foo`
2. Replace the `new_http_archive` calls with the commented out ones below them in WORKSPACE
3. run `bazel build //foo`

You'll notice that the second build occurs quickly and will have built the
target against the version of Go that was commented out. The new version was not downloaded.
