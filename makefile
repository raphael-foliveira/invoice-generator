run: build
	./dist/invgen . -rod=show

build-all: build build-windows build-macos

build: 
	go build -o dist/invgen cmd/main.go

build-windows:
	GOOS=windows go build -o dist/invgen.exe cmd/main.go

build-macos:
	GOOS=darwin go build -o dist/invgen cmd/main.go