bootstrap:
	glide install
	glide update --strip-vendor

build:
	cd cmd/compass && go build
test:
	go test -cover $(shell go list ./... | grep -v /vendor/)
clean:
	rm cmd/compass/compass
