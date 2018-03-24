.PHONY: test
test:
	go test -v -cover ./rcl

.PHONY: fmt
fmt:
	gofmt -d -l -w ./rcl

.PHONY: install
install:
	go install ./cmd/rcl

.PHONY: dep-install
dep-install:
	dep ensure