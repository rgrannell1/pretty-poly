


setGoPath:
	export GOPATH="/home/ryan/Code/pretty_poly"

build: setGoPath

	go build github.com/rgrannell/pretty_poly
	go install github.com/rgrannell/pretty_poly

test: setGoPath

	go build github.com/rgrannell/pretty_poly
	go install github.com/rgrannell/pretty_poly
	go test github.com/rgrannell/pretty_poly -v

bench: setGoPath

	go build github.com/rgrannell/pretty_poly
	go install github.com/rgrannell/pretty_poly

	go test github.com/rgrannell/pretty_poly -v -bench . -cover -benchmem -test.run Benchmark
