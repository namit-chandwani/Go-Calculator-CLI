.DEFAULT_GOAL := build_and_run
BIN_FILE=calculator

build:
	@go build -o "${BIN_FILE}" ./cmd/calculator
run:
	./"${BIN_FILE}"
build_and_run: build run
clean:
	go clean
	rm -f ${BIN_FILE}