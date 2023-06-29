BINARY_FILE = project_euler_go

build: 
	go build .

PROBLEM ?= 1
INPUT ?= 1

run: build
	./${BINARY_FILE} --problem ${PROBLEM} --input ${INPUT}

clean:
	go clean

fmt:
	go fmt ./...

vet:
	go vet ./...
