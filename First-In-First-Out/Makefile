BIN := fifo

all: test $(BIN)

test: main.go main_test.go
	go mod tidy
	go test -v ./...

$(BIN): main.go
	go build -o $(BIN) ./...