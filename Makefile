.PHONY: tool.install
tool.install:
	@go install github.com/golang/mock/mockgen@latest
	@go install github.com/bazelbuild/bazelisk@latest

.PHONY: mock.gen
mock.gen:
	@mockgen -source interfaces.go -destination mock/mock.go

.PHONY: bazel.update
bazel.update:
	@bazelisk run //:gazelle
	@bazelisk run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
	@bazelisk run //:gazelle

.PHONY: bazel.test
bazel.test:
	@bazelisk test --test_output=all //...

.PHONY: bazel.build
bazel.build:
	@bazelisk build //cmd

