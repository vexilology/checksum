default_build := go build hashgosum.go

all: install build tests

list:
	@grep '^[^#[:space:]].*:' Makefile

build:
	$(default_build)

linux32:
	GOOS=linux GOARCH=386 $(default_build)

linux64:
	GOOS=linux GOARCH=amd64 $(default_build)

win32:
	GOOS=windows GOARCH=386 $(default_build)

win64:
	GOOS=windows GOARCH=amd64 $(default_build)

macos32:
	GOOS=darwin GOARCH=386 $(default_build)

macos64:
	GOOS=darwin GOARCH=amd64 $(default_build)

install:
	go get github.com/cxmcc/tiger
	go get github.com/htruong/go-md2
	go get github.com/ebfe/keccak
	go get -u golang.org/x/crypto

fix:
	go get -u golang.org/x/sys/cpu

tests:
	cd algorithm && go test -bench=.
	cd algorithm && go test -v
