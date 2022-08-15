mocks:
	rm -fr mock && \
	mkdir mock && \
	mockgen -source=http/health_check.go -destination=mock/mock_health_check.go -package=mock

test:
	go test ./... -covermode=count -v -coverprofile cp.out

build:
	rm -rf bin && \
	go mod download && \
	go mod vendor && \
	go build -o bin/go-utils