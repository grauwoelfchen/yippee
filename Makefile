proj := "github.com/grauwoelfchen/yippee"

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
