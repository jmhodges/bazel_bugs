To test this bug run:

1. run `bazel test //foo/...` in Travis CI (or maybe any linux machine?)

You'll notice the tests fail but work on macOS (or on maybe a non-containerized linux machine)
