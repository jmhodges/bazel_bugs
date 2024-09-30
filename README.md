Run `mkdir -p .vscode && cp vscode-settings.json .vscode/settings.json`

Then run `bazel build //foo/...`

Then open this directory as a workspace. Then open foo/foo.go in that workspace.

Wait for about 10 seconds or so.

Watch the lower bar for an error notification (and, on my machine, for the
./bazel_bugs/external/ directory to disappear).
