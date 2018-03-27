NV = $(shell glide novendor)
.PHONY: all clean build linux vet lint test

default: build

clean:
	rm -rf bin/application

dep:
	glide install -v

build: clean dep
	go build -v -o application

linux: clean dep
	env GOOS=linux GOARCH=amd64 go build -v

vet:
	go vet $(NV)

lint:
	golint $(NV)

test: vet lint
	go test $(NV) -v -race