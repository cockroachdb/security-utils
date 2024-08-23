# generate-bazel regenerates all BUILD.bazel files within security-utils.
.PHONY: generate-bazel
generate-bazel:
	bazel run //:gazelle
