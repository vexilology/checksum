help:
	@echo "make install (install all packages)"
	@echo "make build (build project)"
	@echo "make fix (if there is an error related to the 'sys/cpu' library)"

build:
	@echo "Building..."
	go build -o hashgosum hashgosum.go

install:
	@echo "Installing packages..."
	go get github.com/cxmcc/tiger
	go get github.com/htruong/go-md2
	go get github.com/ebfe/keccak
	go get -u golang.org/x/crypto

fix:
	go get -u golang.org/x/sys/cpu

test:
	go test -v
