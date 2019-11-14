package main

import (
	"fmt"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"io/ioutil"
	"os"
	"log"
)

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err !=  nil {
		log.Fatal(err)
	}

	fmt.Printf("MD5: %x\n\n", md5.Sum(data))
	fmt.Printf("SHA1: %x\n\n", sha1.Sum(data))
	fmt.Printf("SHA256: %x\n\n", sha256.Sum256(data))
	fmt.Printf("SHA512: %x\n\n", sha512.Sum512(data))
}