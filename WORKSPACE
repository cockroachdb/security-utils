# Define the top level namespace. This lets everything be addressable using
# `@com_github_cockroachdb_cockroach//...`.
workspace(name = "com_github_cockroachdb_security_utils")

# Load the things that let us load other things.
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Load go bazel tools. This gives us access to the go bazel SDK/toolchains.
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "ada68324bc20ffd1b557bab4cf8dba9b742570a46a505b0bc99c1fde5132cce5",
    strip_prefix = "cockroachdb-rules_go-734c37d",
    urls = [
        # cockroachdb/rules_go as of 734c37d6e2e6570f420d27fd97424b2fe2b402af
        # (upstream release-0.46 plus a few patches).
        "https://storage.googleapis.com/public-bazel-artifacts/bazel/cockroachdb-rules_go-v0.27.0-459-g734c37d.tar.gz",
    ],
)

# Load gazelle. This lets us auto-generate BUILD.bazel files throughout the
# repo.
http_archive(
    name = "bazel_gazelle",
    sha256 = "22140e6a7a28df5ec7477f12b286f24dedf8dbef0a12ffbbac10ae80441aa093",
    strip_prefix = "bazelbuild-bazel-gazelle-061cc37",
    urls = [
        "https://storage.googleapis.com/public-bazel-artifacts/bazel/bazelbuild-bazel-gazelle-v0.33.0-0-g061cc37.zip",
    ],
)

# Load the go dependencies and invoke them.
load(
    "@io_bazel_rules_go//go:deps.bzl",
    "go_register_toolchains",
    "go_rules_dependencies",
)

# Load gazelle dependencies.
load(
    "@bazel_gazelle//:deps.bzl",
    "gazelle_dependencies",
    "go_repository",
)

go_rules_dependencies()

go_register_toolchains(version = "1.20.5")

gazelle_dependencies()