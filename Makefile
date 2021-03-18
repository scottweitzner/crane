DIST_DIR = $(shell pwd)/dist

.PHONY: clean crane crane-linux

clean:
	rm -rf ${DIST_DIR}

crane:
	@go build -o ${DIST_DIR}/crane main.go

crane-linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make crane
