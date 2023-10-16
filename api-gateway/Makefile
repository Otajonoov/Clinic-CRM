CURRENT_DIR=$(shell pwd)
APP=template
APP_CMD_DIR=./cmd


run:
	go run cmd/main.go

proto-gen:
	./scripts/gen-proto.sh

pull-sub-module:
	git submodule update --init --recursive

update-sub-module:
	git submodule update --remote --merge

swag-gen:
	swag init -g api/router.go -o api/docs