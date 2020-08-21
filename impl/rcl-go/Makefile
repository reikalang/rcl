VERSION = 0.0.1
BUILD_COMMIT = $(shell git rev-parse HEAD)
BUILD_TIME = $(shell date +%Y-%m-%dT%H:%M:%S%z)
CURRENT_USER = $(USER)
FLAGS = -X main.version=$(VERSION) -X main.commit=$(BUILD_COMMIT) -X main.buildTime=$(BUILD_TIME) -X main.buildUser=$(CURRENT_USER)
PKG=github.com/reikalang/rcl/impl/rcl-go

.PHONY: test
test:
	go test -v -cover ./rcl/...

.PHONY: fmt
fmt:
	gommon format -d -l -w ./rcl ./cmd

.PHONY: install
install:
	go install -ldflags "$(FLAGS)" ./cmd/rcl

.PHONY: doc
doc:
	xdg-open http://localhost:6060/pkg/$(PKG) &
	godoc -http=":6060"

.PHONY: package
package:
	rm -f rcl-v$(VERSION)-linux-amd64.zip
	go build -ldflags "$(FLAGS)" -o rcl-v$(VERSION)-linux-amd64 ./cmd/rcl
	zip rcl-v$(VERSION)-linux-amd64.zip rcl-v$(VERSION)-linux-amd64
	rm rcl-v$(VERSION)-linux-amd64

# --- docker --- #
.PHONY: docker-test
docker-test:
	docker-compose run --rm go19
	docker-compose run --rm go110
	docker-compose run --rm golatest
# --- docker --- #