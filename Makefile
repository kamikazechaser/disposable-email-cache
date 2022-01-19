BIN := disposable-email-cache
GIN := jsoniter

.PHONY: build
build:
	CGO_ENABLED=1 GOOS=linux go build -tags=${GIN} -o ${BIN} -ldflags="-s -w" cmd/*/*.go