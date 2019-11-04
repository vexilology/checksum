package main

import (
	"fmt"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

func main() {

	privatemessage := "test sum"

	MD5 := fmt.Sprintf("%x", md5.Sum([]byte(privatemessage)))
	SHA1 := fmt.Sprintf("%x", sha1.Sum([]byte(privatemessage)))
	SHA256 := fmt.Sprintf("%x", sha256.Sum256([]byte(privatemessage)))
	SHA512 := fmt.Sprintf("%x", sha512.Sum512([]byte(privatemessage)))

	fmt.Println("MD5: ", MD5)
	fmt.Println("SHA-1: ", SHA1)
	fmt.Println("SHA-256: ", SHA256)
	fmt.Println("SHA-512: ", SHA512)
}
