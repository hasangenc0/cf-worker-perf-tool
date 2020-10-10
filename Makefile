BINARY_NAME=cf-perf

all: build
build:
	go mod vendor
	go build -o $(BINARY_NAME) -v
clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	go clean
	go build -o $(BINARY_NAME) -v
	./$(BINARY_NAME)
