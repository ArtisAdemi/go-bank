build: 
	@go build -o bin/proj2

run: build
	@./bin/proj2

test:
	@go test -v ./...