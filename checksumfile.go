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
	"encoding/hex"
)

func main() {
	md5 := md5.New()
	f, err := ioutil.ReadFile(os.Args[1])
	md5.Write(f)
	if err !=  nil {
		log.Fatal(err)
	}

	os.Stdout.WriteString(hex.EncodeToString(md5.Sum(nil)))
	fmt.Println(": <- MD5")

	sha1 := sha1.New()
	e, err := ioutil.ReadFile(os.Args[1])
	sha1.Write(e)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.WriteString(hex.EncodeToString(sha1.Sum(nil)))
	fmt.Println(": <- SHA1")

	sha256 := sha256.New()
	s, err := ioutil.ReadFile(os.Args[1])
	sha256.Write(s)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.WriteString(hex.EncodeToString(sha256.Sum(nil)))
	fmt.Println(": <- SHA256")

	sha512 := sha512.New()
	x, err := ioutil.ReadFile(os.Args[1])
	sha512.Write(x)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.WriteString(hex.EncodeToString(sha512.Sum(nil)))
	fmt.Println(": <- SHA512")
}