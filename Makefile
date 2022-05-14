run: build
	./bin/prunus

build: build-main
	mkdir -p bin

build-main:
	cd pkg && go build -o ../bin/prunus