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

install: bactl
	sudo cp bin/bactl /usr/local/bin/bactl
	sudo chmod +x /usr/local/bin/bactl

uninstall:
	sudo rm -f /usr/local/bin/bactl

bin:
	mkdir -p bin
