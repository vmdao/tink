licenses(["notice"])  # Apache 2.0

cc_library(
    name = "hybrid_decrypt_set_wrapper",
    srcs = ["hybrid_decrypt_set_wrapper.cc"],
    hdrs = ["hybrid_decrypt_set_wrapper.h"],
    visibility = ["//visibility:private"],
    deps = [
        "//cc:crypto_format",
        "//cc:hybrid_decrypt",
        "//cc:primitive_set",
        "//cc/util:status",
        "//cc/util:statusor",
        "//proto:tink_cc_proto",
        "@com_google_protobuf//:protobuf_lite",
    ],
)

cc_library(
    name = "hybrid_encrypt_set_wrapper",
    srcs = ["hybrid_encrypt_set_wrapper.cc"],
    hdrs = ["hybrid_encrypt_set_wrapper.h"],
    visibility = ["//visibility:private"],
    deps = [
        "//cc:crypto_format",
        "//cc:hybrid_encrypt",
        "//cc:primitive_set",
        "//cc/util:status",
        "//cc/util:statusor",
        "//proto:tink_cc_proto",
        "@com_google_protobuf//:protobuf_lite",
    ],
)

cc_library(
    name = "hybrid_decrypt_catalogue",
    srcs = ["hybrid_decrypt_catalogue.cc"],
    hdrs = ["hybrid_decrypt_catalogue.h"],
    visibility = ["//visibility:public"],
    deps = [
        ":ecies_aead_hkdf_private_key_manager",
        "//cc:catalogue",
        "//cc:hybrid_decrypt",
        "//cc:key_manager",
        "//cc/util:status",
        "//cc/util:statusor",
        "//cc/util:strings",
    ],
)

cc_library(
    name = "hybrid_decrypt_config",
    srcs = ["hybrid_decrypt_config.cc"],
    hdrs = ["hybrid_decrypt_config.h"],
    visibility = ["//visibility:public"],
    deps = [
        ":hybrid_decrypt_catalogue",
        "//cc:hybrid_decrypt",
        "//cc/aead:aead_config",
        "//cc/util:status",
    ],
)

cc_library(
    name = "hybrid_decrypt_factory",
    srcs = ["hybrid_decrypt_factory.cc"],
    hdrs = ["hybrid_decrypt_factory.h"],
    visibility = ["//visibility:public"],
    deps = [
        ":hybrid_decrypt_set_wrapper",
        "//cc:hybrid_decrypt",
        "//cc:key_manager",
        "//cc:keyset_handle",
        "//cc:primitive_set",
        "//cc:registry",
        "//cc/util:status",
        "//cc/util:statusor",
    ],
)

cc_library(
    name = "hybrid_encrypt_catalogue",
    srcs = ["hybrid_encrypt_catalogue.cc"],
    hdrs = ["hybrid_encrypt_catalogue.h"],
    visibility = ["//visibility:public"],
    deps = [
        ":ecies_aead_hkdf_public_key_manager",
        "//cc:catalogue",
        "//cc:hybrid_encrypt",
        "//cc:key_manager",
        "//cc/util:status",
        "//cc/util:statusor",
        "//cc/util:strings",
    ],
)

cc_library(
    name = "hybrid_encrypt_config",
    srcs = ["hybrid_encrypt_config.cc"],
    hdrs = ["hybrid_encrypt_config.h"],
    visibility = ["//visibility:public"],
    deps = [
        ":hybrid_encrypt_catalogue",
        "//cc:hybrid_encrypt",
        "//cc/aead:aead_config",
        "//cc/util:status",
    ],
)

cc_library(
    name = "hybrid_encrypt_factory",
    srcs = ["hybrid_encrypt_factory.cc"],
    hdrs = ["hybrid_encrypt_factory.h"],
    visibility = ["//visibility:public"],
    deps = [
        ":hybrid_encrypt_set_wrapper",
        "//cc:hybrid_encrypt",
        "//cc:key_manager",
        "//cc:keyset_handle",
        "//cc:primitive_set",
        "//cc:registry",
        "//cc/util:status",
        "//cc/util:statusor",
    ],
)

cc_library(
    name = "ecies_aead_hkdf_dem_helper",
    srcs = ["ecies_aead_hkdf_dem_helper.cc"],
    hdrs = ["ecies_aead_hkdf_dem_helper.h"],
    visibility = ["//visibility:private"],
    deps = [
        "//cc:aead",
        "//cc:key_manager",
        "//cc:registry",
        "//cc/util:ptr_util",
        "//cc/util:status",
        "//cc/util:statusor",
        "//proto:aes_gcm_cc_proto",
        "//proto:common_cc_proto",
        "//proto:tink_cc_proto",
    ],
)

cc_library(
    name = "ecies_aead_hkdf_hybrid_decrypt",
    srcs = ["ecies_aead_hkdf_hybrid_decrypt.cc"],
    hdrs = ["ecies_aead_hkdf_hybrid_decrypt.h"],
    visibility = ["//visibility:private"],
    deps = [
        ":ecies_aead_hkdf_dem_helper",
        "//cc:hybrid_decrypt",
        "//cc/subtle:ec_util",
        "//cc/subtle:ecies_hkdf_recipient_kem_boringssl",
        "//cc/util:ptr_util",
        "//cc/util:status",
        "//cc/util:statusor",
        "//proto:common_cc_proto",
        "//proto:ecies_aead_hkdf_cc_proto",
        "//proto:tink_cc_proto",
    ],
)

cc_library(
    name = "ecies_aead_hkdf_hybrid_encrypt",
    srcs = ["ecies_aead_hkdf_hybrid_encrypt.cc"],
    hdrs = ["ecies_aead_hkdf_hybrid_encrypt.h"],
    visibility = ["//visibility:private"],
    deps = [
        ":ecies_aead_hkdf_dem_helper",
        "//cc:aead",
        "//cc:hybrid_encrypt",
        "//cc:key_manager",
        "//cc:registry",
        "//cc/subtle:ecies_hkdf_sender_kem_boringssl",
        "//cc/util:status",
        "//cc/util:statusor",
        "//proto:common_cc_proto",
        "//proto:ecies_aead_hkdf_cc_proto",
        "//proto:tink_cc_proto",
    ],
)

cc_library(
    name = "ecies_aead_hkdf_private_key_manager",
    srcs = ["ecies_aead_hkdf_private_key_manager.cc"],
    hdrs = ["ecies_aead_hkdf_private_key_manager.h"],
    visibility = ["//objc/hybrid:__pkg__"],
    deps = [
        ":ecies_aead_hkdf_hybrid_decrypt",
        "//cc:hybrid_decrypt",
        "//cc:key_manager",
        "//cc:registry",
        "//cc/util:status",
        "//cc/util:statusor",
        "//proto:common_cc_proto",
        "//proto:ecies_aead_hkdf_cc_proto",
        "//proto:tink_cc_proto",
    ],
)

cc_library(
    name = "ecies_aead_hkdf_public_key_manager",
    srcs = ["ecies_aead_hkdf_public_key_manager.cc"],
    hdrs = ["ecies_aead_hkdf_public_key_manager.h"],
    visibility = ["//objc/hybrid:__pkg__"],
    deps = [
        ":ecies_aead_hkdf_hybrid_encrypt",
        "//cc:hybrid_encrypt",
        "//cc:key_manager",
        "//cc:registry",
        "//cc/util:status",
        "//cc/util:statusor",
        "//proto:common_cc_proto",
        "//proto:ecies_aead_hkdf_cc_proto",
        "//proto:tink_cc_proto",
    ],
)

# tests

cc_test(
    name = "hybrid_decrypt_set_wrapper_test",
    size = "small",
    srcs = ["hybrid_decrypt_set_wrapper_test.cc"],
    copts = ["-Iexternal/gtest/include"],
    deps = [
        ":hybrid_decrypt_set_wrapper",
        "//cc:hybrid_decrypt",
        "//cc:primitive_set",
        "//cc/util:status",
        "//cc/util:test_util",
        "//proto:tink_cc_proto",
        "@com_google_googletest//:gtest_main",
    ],
)

cc_test(
    name = "hybrid_encrypt_set_wrapper_test",
    size = "small",
    srcs = ["hybrid_encrypt_set_wrapper_test.cc"],
    copts = ["-Iexternal/gtest/include"],
    deps = [
        ":hybrid_encrypt_set_wrapper",
        "//cc:hybrid_encrypt",
        "//cc:primitive_set",
        "//cc/util:status",
        "//cc/util:test_util",
        "//proto:tink_cc_proto",
        "@com_google_googletest//:gtest_main",
    ],
)

cc_test(
    name = "hybrid_decrypt_catalogue_test",
    size = "small",
    srcs = ["hybrid_decrypt_catalogue_test.cc"],
    copts = ["-Iexternal/gtest/include"],
    deps = [
        ":hybrid_decrypt_catalogue",
        ":hybrid_decrypt_config",
        "//cc:catalogue",
        "//cc:hybrid_decrypt",
        "//cc/util:status",
        "//cc/util:statusor",
        "@com_google_googletest//:gtest_main",
    ],
)

cc_test(
    name = "hybrid_decrypt_config_test",
    size = "small",
    srcs = ["hybrid_decrypt_config_test.cc"],
    copts = ["-Iexternal/gtest/include"],
    deps = [
        ":hybrid_decrypt_config",
        "//cc:catalogue",
        "//cc:config",
        "//cc:hybrid_decrypt",
        "//cc:registry",
        "//cc/util:status",
        "@com_google_googletest//:gtest_main",
    ],
)

cc_test(
    name = "hybrid_decrypt_factory_test",
    size = "small",
    srcs = ["hybrid_decrypt_factory_test.cc"],
    copts = ["-Iexternal/gtest/include"],
    deps = [
        ":ecies_aead_hkdf_public_key_manager",
        ":hybrid_decrypt_config",
        ":hybrid_decrypt_factory",
        "//cc:hybrid_decrypt",
        "//cc:hybrid_encrypt",
        "//cc:keyset_handle",
        "//cc/util:ptr_util",
        "//cc/util:status",
        "//cc/util:test_util",
        "//proto:ecies_aead_hkdf_cc_proto",
        "//proto:tink_cc_proto",
        "@com_google_googletest//:gtest_main",
    ],
)

cc_test(
    name = "hybrid_encrypt_catalogue_test",
    size = "small",
    srcs = ["hybrid_encrypt_catalogue_test.cc"],
    copts = ["-Iexternal/gtest/include"],
    deps = [
        ":hybrid_encrypt_catalogue",
        ":hybrid_encrypt_config",
        "//cc:catalogue",
        "//cc:hybrid_encrypt",
        "//cc/util:status",
        "//cc/util:statusor",
        "@com_google_googletest//:gtest_main",
    ],
)

cc_test(
    name = "hybrid_encrypt_config_test",
    size = "small",
    srcs = ["hybrid_encrypt_config_test.cc"],
    copts = ["-Iexternal/gtest/include"],
    deps = [
        ":hybrid_encrypt_config",
        "//cc:catalogue",
        "//cc:config",
        "//cc:hybrid_encrypt",
        "//cc:registry",
        "//cc/util:status",
        "@com_google_googletest//:gtest_main",
    ],
)

cc_test(
    name = "hybrid_encrypt_factory_test",
    size = "small",
    srcs = ["hybrid_encrypt_factory_test.cc"],
    copts = ["-Iexternal/gtest/include"],
    deps = [
        ":hybrid_encrypt_config",
        ":hybrid_encrypt_factory",
        "//cc:hybrid_encrypt",
        "//cc:keyset_handle",
        "//cc/util:status",
        "//cc/util:test_util",
        "//proto:ecies_aead_hkdf_cc_proto",
        "//proto:tink_cc_proto",
        "@com_google_googletest//:gtest_main",
    ],
)

cc_test(
    name = "ecies_aead_hkdf_hybrid_decrypt_test",
    size = "small",
    srcs = ["ecies_aead_hkdf_hybrid_decrypt_test.cc"],
    copts = ["-Iexternal/gtest/include"],
    deps = [
        ":ecies_aead_hkdf_hybrid_decrypt",
        ":ecies_aead_hkdf_hybrid_encrypt",
        "//cc:hybrid_decrypt",
        "//cc/aead:aes_gcm_key_manager",
        "//cc/subtle:random",
        "//cc/subtle:subtle_util_boringssl",
        "//cc/util:ptr_util",
        "//cc/util:status",
        "//cc/util:statusor",
        "//cc/util:test_util",
        "//proto:common_cc_proto",
        "//proto:ecies_aead_hkdf_cc_proto",
        "@com_google_googletest//:gtest_main",
    ],
)

cc_test(
    name = "ecies_aead_hkdf_hybrid_encrypt_test",
    size = "small",
    srcs = ["ecies_aead_hkdf_hybrid_encrypt_test.cc"],
    copts = ["-Iexternal/gtest/include"],
    deps = [
        ":ecies_aead_hkdf_hybrid_encrypt",
        "//cc:hybrid_encrypt",
        "//cc:registry",
        "//cc/aead:aes_gcm_key_manager",
        "//cc/subtle:random",
        "//cc/subtle:subtle_util_boringssl",
        "//cc/util:ptr_util",
        "//cc/util:status",
        "//cc/util:statusor",
        "//cc/util:test_util",
        "//proto:common_cc_proto",
        "//proto:ecies_aead_hkdf_cc_proto",
        "@com_google_googletest//:gtest_main",
    ],
)

cc_test(
    name = "ecies_aead_hkdf_private_key_manager_test",
    size = "small",
    srcs = ["ecies_aead_hkdf_private_key_manager_test.cc"],
    copts = ["-Iexternal/gtest/include"],
    deps = [
        ":ecies_aead_hkdf_private_key_manager",
        ":ecies_aead_hkdf_public_key_manager",
        "//cc:hybrid_decrypt",
        "//cc/aead:aes_gcm_key_manager",
        "//cc/util:status",
        "//cc/util:statusor",
        "//cc/util:test_util",
        "//proto:aes_eax_cc_proto",
        "//proto:common_cc_proto",
        "//proto:ecies_aead_hkdf_cc_proto",
        "//proto:tink_cc_proto",
        "@com_google_googletest//:gtest_main",
    ],
)

cc_test(
    name = "ecies_aead_hkdf_public_key_manager_test",
    size = "small",
    srcs = ["ecies_aead_hkdf_public_key_manager_test.cc"],
    copts = ["-Iexternal/gtest/include"],
    deps = [
        ":ecies_aead_hkdf_public_key_manager",
        "//cc:hybrid_encrypt",
        "//cc:registry",
        "//cc/aead:aes_gcm_key_manager",
        "//cc/util:status",
        "//cc/util:statusor",
        "//cc/util:test_util",
        "//proto:aes_eax_cc_proto",
        "//proto:common_cc_proto",
        "//proto:ecies_aead_hkdf_cc_proto",
        "//proto:tink_cc_proto",
        "@com_google_googletest//:gtest_main",
    ],
)
