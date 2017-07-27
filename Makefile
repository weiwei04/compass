bootstrap:
	glide install

build:
	cd cmd/compass && go build
	cd cmd/helm-registry-plugin && go build
test:
	go test -cover $(shell go list ./... | grep -v /vendor/)
clean:
	rm cmd/compass/compass
	rm cmd/helm-registry-plugin/helm-registry-plugin
