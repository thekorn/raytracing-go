BINARY_NAME=main.out
OUT_DIR=tmp
 
all: build test
 
outdir:
	mkdir -p ${OUT_DIR}

build:
	go build -o ${BINARY_NAME} main.go
 
test:
	go test -v main.go
 
run: build outdir
	./${BINARY_NAME}
 
clean:
	go clean
	rm ${BINARY_NAME}
	rm -rf ${OUT_DIR}