BINARY_NAME=main.out
OUT_DIR=tmp
 
all: build test
 
outdir:
	mkdir -p ${OUT_DIR}

build:
	go build -o ${BINARY_NAME} main.go
 
test:
	go test -coverprofile cover.out -v ./...

cover: test
	go tool cover -html=cover.out
 
run: build outdir
	./${BINARY_NAME}

open: run
	open ${OUT_DIR}/go.ppm
 
clean:
	go clean
	rm ${BINARY_NAME}
	rm -rf ${OUT_DIR}