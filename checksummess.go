package main

import (
	"strconv"
	"fmt"
	"bufio"
	"os"
	"encoding/base32"
	"encoding/base64"
	"hash/crc32"
	"hash/adler32"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/ripemd160"
)

func main() {

	var data string

	fmt.Println("**** Enter your message ****")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		data = scanner.Text()
	}

	firstCRC32 := crc32.ChecksumIEEE([]byte(data))
	secondCRC32 := strconv.FormatUint(uint64(firstCRC32), 16)
	firstADLER32 := adler32.Checksum([]byte(data))
	secondADLER32 := strconv.FormatUint(uint64(firstADLER32), 16)
	firstMD4 := md4.New()
	firstMD4.Write([]byte(data))
	secondMD4 := fmt.Sprintf("%x", firstMD4.Sum([]byte(nil)))
	MD5 := fmt.Sprintf("%x", md5.Sum([]byte(data)))
	SHA1 := fmt.Sprintf("%x", sha1.Sum([]byte(data)))
	SHA256 := fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
	SHA384 := fmt.Sprintf("%x", sha512.Sum384([]byte(data)))
	SHA512 := fmt.Sprintf("%x", sha512.Sum512([]byte(data)))
	BASE32 := base32.StdEncoding.EncodeToString([]byte(data))
	BASE64 := base64.StdEncoding.EncodeToString([]byte(data))
	BLAKE2B256 := fmt.Sprintf("%x", blake2b.Sum256([]byte(data)))
	BLAKE2B512 := fmt.Sprintf("%x", blake2b.Sum512([]byte(data)))
	firstRIPEMD160 := ripemd160.New()
	firstRIPEMD160.Write([]byte(data))
	secondRIPEMD160 := fmt.Sprintf("%x", firstRIPEMD160.Sum(nil))

	fmt.Print("-----------------------\n")
	fmt.Println("BASE32:", BASE32)
	fmt.Println("===")
	fmt.Println("BASE64:", BASE64)
	fmt.Println("===")
	fmt.Println("CRC32:", secondCRC32)
	fmt.Println("===")
	fmt.Println("ADLER32:", secondADLER32)
	fmt.Println("===")
	fmt.Println("BLAKE2B-256:", BLAKE2B256)
	fmt.Println("===")
	fmt.Println("BLAKE2B-512:", BLAKE2B512)
	fmt.Println("===")
	fmt.Println("RIPEMD160:", secondRIPEMD160)
	fmt.Println("===")
	fmt.Println("MD4:", secondMD4)
	fmt.Println("===")
	fmt.Println("MD5:", MD5)
	fmt.Println("===")
	fmt.Println("SHA1:", SHA1)
	fmt.Println("===")
	fmt.Println("SHA256:", SHA256)
	fmt.Println("===")
	fmt.Println("SHA384:", SHA384)
	fmt.Println("===")
	fmt.Println("SHA512:", SHA512)
}