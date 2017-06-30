
CURRENT_DIR = $(shell pwd)





setGoPath:
	export GOPATH=$(CURRENT_DIR)

fetchDependencies: setGoPath
	go get github.com/gonum

build: setGoPath

	go build github.com/rgrannell/pretty_poly
	go install github.com/rgrannell/pretty_poly

test: FORCE

	docker build -t pretty_poly_tests -f dockerfiles/test-pretty-poly.txt .
	docker run -t pretty_poly_tests

bench: force-dangerous

	docker build -t pretty_poly_benchmarks -f dockerfiles/bench-pretty-poly.txt .
	docker run -t pretty_poly_benchmarks

vet: setGoPath
	go vet github.com/rgrannell/pretty_poly

install: snap
	cd snapcraft && snap install pretty_poly_* --force-dangerous && cd ..

snap: FORCE
	cd snapcraft && snapcraft clean && snapcraft snap && cd ..

run: FORCE

	docker build -t pretty_poly -f dockerfiles/pretty-poly.txt .
	docker run -t pretty_poly

run-debug: FORCE

	docker run -i -t pretty_poly /bin/bash

FORCE:
