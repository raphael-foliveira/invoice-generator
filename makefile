run: build
	./dist/main -rod=show

build-all: build build-windows

build: 
	go build -o dist/main cmd/main.go

build-windows:
	GOOS=windows go build -o dist/main.exe cmd/main.go