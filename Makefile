build:
	go build -o bin/gex
run: build
	./bin/gex
test:
	go test -v ./...
