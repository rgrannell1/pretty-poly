
CURRENT_DIR = $(shell pwd)





setGoPath:
	export GOPATH=$(CURRENT_DIR)

fetchDependencies: setGoPath
	go get github.com/gonum

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

	go test github.com/rgrannell/pretty_poly -v -bench . -cover -benchmem -test.run Benchmark -count 2

vet: setGoPath
	go vet github.com/rgrannell/pretty_poly

install: snap
	cd snapcraft && snap install pretty_poly_* --force-dangerous && cd ..

snap: FORCE
	cd snapcraft && snapcraft clean && snapcraft snap && cd ..

FORCE:
