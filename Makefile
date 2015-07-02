# To install make in boot2docker
# tce-load -wi make
default: test

test:
	docker run -v \
		"$(shell pwd)":/go/src/github.com/bernos/go-restapi \
		-w /go/src/github.com/bernos/go-restapi \
		golang:1.3 \
		sh -c 'go get -v -t ./... && go test ./...'

build:
	docker run -v \
		"$(shell pwd)":/go/src/github.com/bernos/go-restapi \
		-w /go/src/github.com/bernos/go-restapi \
		golang:1.3 \
		sh -c 'go get -v ./... && \
		CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o main'

run: build  
	docker run \
		--rm \
		-p 8080:8080 \
		-v "$(shell pwd):/opt" \
		tianon/true /opt/main