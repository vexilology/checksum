help:
	@echo "make install (install all packages)"
	@echo "make build (build project)"

build:
	@echo "Building..."
	go build -o hashgosum main.go

install:
	@echo "Installing packages..."
	go get github.com/cxmcc/tiger
	go get github.com/htruong/go-md2
	go get github.com/ebfe/keccak
	go get -u golang.org/x/crypto
