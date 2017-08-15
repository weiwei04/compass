bootstrap:
	glide install
	glide update --strip-vendor

build-compass:
	cd cmd/compass && go build

build-fusion:
	cd cmd/fusion && go build

build: build-compass build-fusion

install-fusion: build-fusion
	rm -rf ~/.helm/plugins/fusion
	mkdir -p ~/.helm/plugins/fusion
	cp _plugin/* ~/.helm/plugins/fusion/
	cp cmd/fusion/fusion ~/.helm/plugins/fusion

test:
	go test -cover $(shell go list ./... | grep -v /vendor/)

clean:
	rm -f cmd/compass/compass
	rm -f cmd/fusion/fusion
