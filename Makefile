.PHONY: generate help mock
.DEFAULT_GOAL := help

generate:
	protoc --go_out=./gen --go-grpc_out=./gen ./*.proto

help:	## https://postd.cc/auto-documented-makefile/
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
