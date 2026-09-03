test:
	go test ./... -covermode=count -v -coverprofile cp.out

build:
	rm -rf bin && \
	go mod download && \
	go mod vendor && \
	go build -o bin/go-utils