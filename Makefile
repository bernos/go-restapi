# To install make in boot2docker
# tce-load -wi make

# Name of the binary to be output. On windows the .exe suffix will be added
NAME=main

# Output directory
OUTPUT_DIR=./dist

# Path to the compiled binary
OUTPUT=$(OUTPUT_DIR)/$(NAME)

ifeq ($(OS),Windows_NT)
	# Force cmd.exe as shell on windows to relieve
	# Interrupt/Exception caught (code = 0xc00000fd, addr = 0x4227d3)
	# See http://superuser.com/questions/375029/make-interrupt-exception-caught
	SHELL=C:/Windows/System32/cmd.exe
	OUTPUT=$(OUTPUT_DIR)/$(NAME).exe
endif

all: build

clean:
	go clean -i -x ./...
	-rm -rf $(OUTPUT_DIR)

deps:
	go get -v github.com/tools/godep && \
	$(GOPATH)/bin/godep restore ./... 

test: deps
	$(GOPATH)/bin/godep go test -v ./...

compile: test
	go build -o $(OUTPUT)

copy:
	cp -r ./config $(OUTPUT_DIR)/

build: compile copy

run: all
	$(OUTPUT)

docker-test:
	docker run --dns=54.252.183.4 -v \
		"$(shell pwd)":/go/src/github.com/bernos/go-restapi \
		-w /go/src/github.com/bernos/go-restapi \
		golang:1.3 \
		sh -c 'go get -v github.com/tools/godep &&
			godep restore ./... && \
			godep go test -v ./...'

		#sh -c 'go get -v -t ./... && go test ./...'

docker-build:
	docker run --dns=54.252.183.4 -v \
		"$(shell pwd)":/go/src/github.com/bernos/go-restapi \
		-w /go/src/github.com/bernos/go-restapi \
		golang:1.3 \
		sh -c 'go get -v github.com/tools/godep && \
			godep restore ./... && \
			godep go test -v ./... && \
			CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o main'

#		sh -c 'go get -v ./... && \
#		CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o main'

docker-run: docker-build
	docker run \
		--rm \
		-p 8080:8080 \
		-v "$(shell pwd):/opt" \
		tianon/true /opt/main

.PHONY: test clean compile copy build run
