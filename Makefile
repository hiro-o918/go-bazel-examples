.PHONY: tool.install
tool.install:
	@go install github.com/golang/mock/mockgen@latest
	@go install github.com/bazelbuild/bazelisk@latest

.PHONY: mock.gen
mock.gen:
	@mockgen -source interfaces.go -destination mock/mock.go
