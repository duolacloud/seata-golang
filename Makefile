GOPATH:=$(shell go env GOPATH)
IMAGE_NAME:=tc-server

.PHONY: run
run:
	go run ./cmd/tc/main.go start \
	  -apollo_namespace application \
    -apollo_address http://apollo-dev.dev.lucfish.com:8080 \
    -apollo_app_id tc-server \
    -apollo_cluster=dev \
    --prometheus_addr=

.PHONY: build
build:
	go build -o build/${IMAGE_NAME} *.go

.PHONY: build-linux
build-linux:
	GOOS="linux" GOARCH="amd64" CGO_ENABLED=0 go build -o build/${IMAGE_NAME}-linux ./cmd/tc/main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build -t ${IMAGE_NAME}:latest .

.PHONY: submodule
submodule:
	git submodule init && git submodule update --remote
