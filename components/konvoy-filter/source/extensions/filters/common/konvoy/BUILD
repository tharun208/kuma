licenses(["notice"])  # Apache 2

load(
    "@envoy//bazel:envoy_build_system.bzl",
    "envoy_cc_library",
    "envoy_package",
)

envoy_package()

envoy_cc_library(
    name = "anonymous_reporter_interface",
    hdrs = ["anonymous_reporter.h"],
    repository = "@envoy",
)

envoy_cc_library(
    name = "anonymous_reporter_lib",
    srcs = ["anonymous_reporter_impl.cc"],
    hdrs = ["anonymous_reporter_impl.h"],
    repository = "@envoy",
    deps = [
        ":anonymous_reporter_interface",
        "@envoy//include/envoy/event:dispatcher_interface",
        "@envoy//include/envoy/network:address_interface",
        "@envoy//include/envoy/network:dns_interface",
        "@envoy//include/envoy/runtime:runtime_interface",
        "@envoy//include/envoy/singleton:instance_interface",
        "@envoy//source/common/common:utility_lib",
        "@envoy//source/common/network:utility_lib",
    ],
)

envoy_cc_library(
    name = "utility_lib",
    srcs = ["utility.cc"],
    hdrs = ["utility.h"],
    repository = "@envoy",
    deps = [
        ":anonymous_reporter_lib",
        "@envoy//include/envoy/server:filter_config_interface",
    ],
)