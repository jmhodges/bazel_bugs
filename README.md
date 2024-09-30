## The GOROOT only repro

```
unset GOROOT # Just in case
$ bazel build //foo/...
$ export GOROOT="$(PWD)/bazel-bazel_bugs/external/rules_go~~go_sdk~bazel_bugs__download_0"
$ bazel build //foo/...
```

and you'll see errors on the second `bazel build //foo/...`


## The original repro using vscode and gopls
Run `mkdir -p .vscode && cp vscode-settings.json .vscode/settings.json`

Then run `bazel build //foo/...`

Then open this directory as a workspace. Then open foo/foo.go in that workspace.

Wait for about 10 seconds or so.

Watch the lower bar for an error notification (and, on my machine, for the
./bazel_bugs/external/ directory to disappear).

This will also happen if you run `bazel build //foo/...` and then hit `Reload Window`
and switch to (or `Reload Window` while on) the foo/foo.go tab