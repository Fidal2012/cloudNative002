build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64

build-m1:
	echo "building httpserver binary"
	mkdir -p bin/arm64
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/arm64
