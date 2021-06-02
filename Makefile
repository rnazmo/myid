.PHONY: mod
mod:
	go get -u ./... && go mod tidy

.PHONY: test
test:
	go test -v ./...

.PHONY: bench
bench:
	go test -bench . -benchmem
