all: test build api run

run:
	@echo "Setting environment variables..." && \
	export HOST="http://http-server" && \
	export PORT=":8080" && \
	export API_ENDPOINT="/network" && \
	./build/api && \
	go run server.go

test:
	@echo "Running tests..."
	go test -v ./... -count=1

build:
	@echo "Building..."
	go build -o build/api .

api: build
	@echo "HTTP-server built successfully."
	
.PHONY: all test build api run
