run: build
	./bin/prunus

build: build-main
	mkdir -p bin

build-main:
	cd src && go build -o ../bin/prunus