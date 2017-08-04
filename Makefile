bootstrap:
	glide install
	glide update --strip-vendor

build:
	cd cmd/compass && go build
	cd cmd/fusion && go build
test:
	go test -cover $(shell go list ./... | grep -v /vendor/)
clean:
	rm cmd/compass/compass
	rm cmd/fusion/fusion
