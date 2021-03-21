all: install build tests

build:
	go build hashgosum.go

linux32:
	GOOS=linux GOARCH=386 go build hashgosum.go

linux64:
	GOOS=linux GOARCH=amd64 go build hashgosum.go

win32:
	GOOS=windows GOARCH=386 go build hashgosum.go

win64:
	GOOS=windows GOARCH=amd64 go build hashgosum.go

macos32:
	GOOS=darwin GOARCH=386 go build hashgosum.go

macos64:
	GOOS=darwin GOARCH=amd64 go build hashgosum.go

install:
	go get github.com/cxmcc/tiger
	go get github.com/htruong/go-md2
	go get github.com/ebfe/keccak
	go get -u golang.org/x/crypto

fix:
	go get -u golang.org/x/sys/cpu

tests:
	go test -v
