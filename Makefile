.PHONY: build build-linux build-darwin build-windows

.PHONE: build
build: build-linux build-darwin build-windows

.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/app

.PHONY: build-darwin
build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/app

.PHONY: build-windows
build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/win/app.exe

.PHONY: clean
clean:
	rm -rf ./bin/

.PHONY: run
run:
	./bin/linux/app