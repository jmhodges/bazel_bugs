To see the bug in action, run

    bazel test //somemain:web_test

Then, to see the fix the bug, swap the `browsers` settings in `web_test` as
described in `somemain/BUILD.bazel` and run that `bazel test` again.

What `custom_browser` and that `chromium-no-sandbox.json` does is add the
`--no-sandbox` flag to the arguments passed to the `chromium` binary. Chromium,
when run in headless mode, seems to require that now after an upgrade from bazel
0.13 to bazel 0.18. Possibly, the problem was an upgrade from macOS High Sierra to Mojave.

The `--no-sandbox` flag is mentioned in Chrome's [original blog post about
headless
mode](https://developers.google.com/web/updates/2017/04/headless-chrome) as a
hack for missing certain environment variables. Not sure what changed in bazel
or macOS to require this.
