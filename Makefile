proj := "gitlab.com/grauwoelfchen/yippee"

vet:
	@go vet ./**/*.go
.PHONY: vet

fmt:
	@goimports -w ./*.go ./cmd/
.PHONY: fmt

lint:
	@golint -set_exit_status main*.go ./cmd/*.go

test:
	@go test $(shell glide novendor) -v parallel 3;
.PHONY: test

build:
	@go build -o yippee ${proj}/cmd
.PHONY: build

clean:
	@git clean -f
.PHONY: clean
