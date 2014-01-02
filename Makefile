GO_DIRS=$(shell find . -type f -name '*.go' | grep -o '.*/' | sort | uniq)
TEST_DIRS=$(shell find . -type f -name '*_test.go' | grep -o '.*/' | sort | uniq)
MAIN_DIRS=$(shell find . -type f -name 'main.go' | grep -o '.*/' | sort | uniq)

all: clean test build

get:
	go get $(GO_DIRS)

test: get
	go test $(TEST_DIRS)

build: get
	go build $(MAIN_DIRS)

clean:
	go clean $(GO_DIRS)
