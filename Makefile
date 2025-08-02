all: build

build: bin gosmc bactl

gosmc: bin
	go build \
	    -ldflags="-s -w" \
	    -o bin/gosmc \
	    cmd/main.go

bactl: bin
	go build \
	    -ldflags="-s -w" \
	    -o bin/bactl \
	    cmd/bactl.go

bin:
	mkdir -p bin
