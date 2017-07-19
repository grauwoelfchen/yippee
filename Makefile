proj := "gitlab.com/grauwoelfchen/yippee"

vet:
	@go vet ./*.go
.PHONY: vet

fmt:
	@goimports -w $(shell glide novendor --no-subdir)
.PHONY: fmt

lint:
	@golint -set_exit_status *.go

test:
	@go test ${proj}/cmd -v parallel 3;
	@go test ${proj} -v parallel 3;
.PHONY: test

build:
	@go build -o yippee ${proj}/cmd
.PHONY: build

clean:
	@git clean -f
.PHONY: clean
