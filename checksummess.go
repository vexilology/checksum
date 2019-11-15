package main

import (
	"strconv"
	"fmt"
	"hash/crc32"
	"hash/adler32"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

func main() {

	var data string

	fmt.Println("Enter the message:")
	fmt.Scan(&data)

	firstCRC32 := crc32.ChecksumIEEE([]byte(data))
	secondCRC32 := strconv.FormatUint(uint64(firstCRC32), 16)
	firstADLER32 := adler32.Checksum([]byte(data))
	secondADLER32 := strconv.FormatUint(uint64(firstADLER32), 16)
	MD5 := fmt.Sprintf("%x", md5.Sum([]byte(data)))
	SHA1 := fmt.Sprintf("%x", sha1.Sum([]byte(data)))
	SHA256 := fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
	SHA384 := fmt.Sprintf("%x", sha512.Sum384([]byte(data)))
	SHA512 := fmt.Sprintf("%x", sha512.Sum512([]byte(data)))

	fmt.Println("CRC32:", secondCRC32)
	fmt.Println("ADLER32:", secondADLER32)
	fmt.Println("MD5:", MD5)
	fmt.Println("SHA1:", SHA1)
	fmt.Println("SHA256:", SHA256)
	fmt.Println("SHA384:", SHA384)
	fmt.Println("SHA512:", SHA512)
}
