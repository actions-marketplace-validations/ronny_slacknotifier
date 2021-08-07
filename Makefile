all: binaries

binaries:
	env CGO_ENABLED=0 go build \
		-a \
		-trimpath \
		-ldflags "-s -w -extldflags '-static'" \
		-installsuffix cgo \
		-tags netgo \
		-o ./bin/ \
		./cmd/...

test: lint vet
	go test -race -cover ./...

lint:
	staticcheck ./...

vet:
	go vet ./...

install:
	go install ./cmd/slack-notify

.PHONY: all binaries test lint vet install
