licenses(["notice"])  # Apache 2.0

package(
    default_visibility = [
        "//java:__subpackages__",
    ],
)

load("//java/build_defs:javac.bzl", "JAVACOPTS")

# public interfaces

java_library(
    name = "protos",
    exports = [
        "//proto:tink_java_proto_lite",
    ],
)

android_library(
    name = "android",
    srcs = glob([
        "*.java",
    ]),
    javacopts = JAVACOPTS,
    deps = [
        ":protos",
        "//java/src/main/java/com/google/crypto/tink:android",
        "//java/src/main/java/com/google/crypto/tink:cleartext_keyset_handle_android",
        "//java/src/main/java/com/google/crypto/tink/subtle",
        "//java/src/main/java/com/google/crypto/tink/subtle:aead",
        "@com_google_code_findbugs_jsr305//jar",
        "@com_google_protobuf_javalite//:protobuf_java_lite",
    ],
)
