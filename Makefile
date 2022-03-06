BINARY_NAME=main.out
OUT_DIR=tmp
 
all: build test
 
outdir:
	mkdir -p ${OUT_DIR}

build:
	go build -o ${BINARY_NAME}

build-profile:
	go build -tags pprof -o ${BINARY_NAME}
 
test:
	go test -coverprofile cover.out -v ./...

cover: test
	go tool cover -html=cover.out
 
run: build outdir
	./${BINARY_NAME}
 
run-profile: build-profile outdir
	./${BINARY_NAME}

profile: run-profile
	go tool pprof -http=":8000" ./${BINARY_NAME} ./${OUT_DIR}/cpu.pprof

open: run
	open ${OUT_DIR}/go.ppm

final: run
	convert ${OUT_DIR}/go.ppm final_result.jpg
 
clean:
	go clean
	rm ${BINARY_NAME}
	rm -rf ${OUT_DIR}