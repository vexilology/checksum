default_build := go build hashgosum.go

all: install tests build

list:
	@grep '^[^#[:space:]].*:' Makefile

build:
	$(default_build)

linux:
	GOOS=linux GOARCH=amd64 $(default_build)

win:
	GOOS=windows GOARCH=amd64 $(default_build)

macos:
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
