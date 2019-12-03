package main

import (
	"fmt"
	b64 "encoding/base64"
	"hash/crc32"
	"hash/adler32"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
        fmt.Println("File not found. Try again --> go run checksumfile.go FILENAME")
        return
    }

	data, err := ioutil.ReadFile(os.Args[1])

	if err !=  nil {
		panic(err)
	}

	firstcrc32 := crc32.ChecksumIEEE([]byte(data))
	secondcrc32 := strconv.FormatUint(uint64(firstcrc32), 16)
	firstadler32 := adler32.Checksum([]byte(data))
	secondadler32 := strconv.FormatUint(uint64(firstadler32), 16)
	xbase64 := b64.URLEncoding.EncodeToString([]byte(data))

	fmt.Print("-----------------------\n")
	fmt.Printf("Length: %d bytes\n", len(data))
	fmt.Println("===")
	fmt.Println("BASE64:", xbase64)
	fmt.Println("===")
	fmt.Println("CRC32:", secondcrc32)
	fmt.Println("===")
	fmt.Println("ADLER32:", secondadler32)
	fmt.Println("===")
	fmt.Printf("MD5: %x\n", md5.Sum(data))
	fmt.Println("===")
	fmt.Printf("SHA1: %x\n", sha1.Sum(data))
	fmt.Println("===")
	fmt.Printf("SHA256: %x\n", sha256.Sum256(data))
	fmt.Println("===")
	fmt.Printf("SHA384: %x\n", sha512.Sum384(data))
	fmt.Println("===")
	fmt.Printf("SHA512: %x\n", sha512.Sum512(data))
}