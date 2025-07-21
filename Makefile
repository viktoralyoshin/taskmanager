APP_NAME = taskmanager
BINARY = bin/${APP_NAME}

build:
	go build -o ${BINARY} cmd/taskmanager/main.go

run: build
	./${BINARY}

test:
	go test -race -cover -v ./...
