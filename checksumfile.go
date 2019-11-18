package main

import (
	"fmt"
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

	fmt.Print("-----------------------\n")
	fmt.Println("CRC32:", secondcrc32)
	fmt.Println("ADLER32:", secondadler32)
	fmt.Printf("MD5: % x\n", md5.Sum(data))
	fmt.Printf("SHA1: % x\n", sha1.Sum(data))
	fmt.Printf("SHA256: % x\n", sha256.Sum256(data))
	fmt.Printf("SHA384: % x\n", sha512.Sum384(data))
	fmt.Printf("SHA512: % x\n", sha512.Sum512(data))
	fmt.Print("-----------------------")
}