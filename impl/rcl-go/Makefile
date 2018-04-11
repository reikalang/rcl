VERSION = 0.0.1
BUILD_COMMIT = $(shell git rev-parse HEAD)
BUILD_TIME = $(shell date +%Y-%m-%dT%H:%M:%S%z)
CURRENT_USER = $(USER)
FLAGS = -X main.version=$(VERSION) -X main.commit=$(BUILD_COMMIT) -X main.buildTime=$(BUILD_TIME) -X main.buildUser=$(CURRENT_USER)

.PHONY: test
test:
	go test -v -cover ./rcl

.PHONY: fmt
fmt:
	gofmt -d -l -w ./rcl

.PHONY: install
install:
	go install -ldflags "$(FLAGS)" ./cmd/rcl

.PHONY: dep-install
dep-install:
	dep ensure

.PHONY: update-dep
update-dep:
	dep ensure -update

.PHONY: package
package:
	rm -f rcl-v$(VERSION)-linux-amd64.zip
	go build -ldflags "$(FLAGS)" -o rcl-v$(VERSION)-linux-amd64 ./cmd/rcl
	zip rcl-v$(VERSION)-linux-amd64.zip rcl-v$(VERSION)-linux-amd64
	rm rcl-v$(VERSION)-linux-amd64