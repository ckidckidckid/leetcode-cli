EXECUTABLE=leetcode_cli
VERSION=$(shell git describe --tags --always --long --dirty)
WINDOWS=$(EXECUTABLE)_windows_amd64_$(VERSION).exe
LINUX=$(EXECUTABLE)_linux_amd64_$(VERSION)
DARWIN=$(EXECUTABLE)_darwin_amd64_$(VERSION)

.PHONY: all test clean

all: test build

test:
	go test ./...

build: windows linux darwin
	@echo version: $(VERSION)

windows: $(WINDOWS)

linux: $(LINUX)

darwin: $(DARWIN)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -v -o bin/$(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)" main.go

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -v -o bin/$(LINUX) -ldflags="-s -w -X main.version=$(VERSION)" main.go

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -v -o bin/$(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)" main.go

clean:
	rm -f $(WINDOWS) $(LINUX) $(DARWIN)

buildlocal:
	go build -o bin/$(EXECUTABLE) main.go