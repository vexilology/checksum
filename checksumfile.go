package main

import (
	"fmt"
	"encoding/base32"
	"encoding/base64"
	"hash/crc32"
	"hash/adler32"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"io/ioutil"
	"os"
	"strconv"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
	"github.com/cxmcc/tiger"
)

func main() {

	if len(os.Args) < 2 {
        fmt.Println("File not found. Try again ---> go run checksumfile.go [FILENAME]")
        return
    }

	data, err := ioutil.ReadFile(os.Args[1])

	if err !=  nil {
		panic(err)
	}

	newkeccak256 := sha3.NewLegacyKeccak256()
	newkeccak256.Write([]byte(data))
	newkeccak512 := sha3.NewLegacyKeccak512()
	newkeccak512.Write([]byte(data))

	firstcrc32 := crc32.ChecksumIEEE([]byte(data))
	secondcrc32 := strconv.FormatUint(uint64(firstcrc32), 16)

	firstadler32 := adler32.Checksum([]byte(data))
	secondadler32 := strconv.FormatUint(uint64(firstadler32), 16)

	xbase32 := base32.StdEncoding.EncodeToString(data)
	xbase64 := base64.StdEncoding.EncodeToString(data)

	xMD4 := md4.New()
	xMD4.Write([]byte(data))

	xRIPEMD160 := ripemd160.New()
	xRIPEMD160.Write([]byte(data))

	xTiger := tiger.New()
	xTiger.Write([]byte(data))

	fmt.Print("-----------------------\n")
	fmt.Printf("Length => %d bytes\n", len(data))
	fmt.Print("-----------------------\n")
	fmt.Println("______Binary______")
	fmt.Print("===\n")
	fmt.Println("BASE32 =>", xbase32)
	fmt.Print("===\n")
	fmt.Println("BASE64 =>", xbase64)
	fmt.Print("===\n")
	fmt.Println("______Hash algorithm______")
	fmt.Print("===\n")
	fmt.Printf("TIGER192,3 => %x\n", xTiger.Sum(nil))
	fmt.Print("===\n")
	fmt.Printf("KECCAK256 => %x\n", newkeccak256.Sum(nil))
	fmt.Print("===\n")
	fmt.Printf("KECCAK512 => %x\n", newkeccak512.Sum(nil))
	fmt.Print("===\n")
	fmt.Println("CRC32 =>", secondcrc32)
	fmt.Print("===\n")
	fmt.Println("ADLER32 =>", secondadler32)
	fmt.Print("===\n")
	fmt.Printf("BLAKE2B-256 => %x\n", blake2b.Sum256(data))
	fmt.Print("===\n")
	fmt.Printf("BLAKE2B-512 => %x\n", blake2b.Sum512(data))
	fmt.Print("===\n")
	fmt.Printf("RIPEMD160 => %x\n", xRIPEMD160.Sum(nil))
	fmt.Print("===\n")
	fmt.Printf("MD4 => %x\n", xMD4.Sum([]byte(nil)))
	fmt.Print("===\n")
	fmt.Printf("MD5 => %x\n", md5.Sum(data))
	fmt.Print("===\n")
	fmt.Printf("SHA1 => %x\n", sha1.Sum(data))
	fmt.Print("===\n")
	fmt.Printf("SHA224 => %x\n", sha256.Sum224(data))
	fmt.Print("===\n")
	fmt.Printf("SHA256 => %x\n", sha256.Sum256(data))
	fmt.Print("===\n")
	fmt.Printf("SHA384 => %x\n", sha512.Sum384(data))
	fmt.Print("===\n")
	fmt.Printf("SHA512 => %x\n", sha512.Sum512(data))
	fmt.Print("===\n")
	fmt.Printf("SHA512_224 => %x\n", sha512.Sum512_224(data))
	fmt.Print("===\n")
	fmt.Printf("SHA512_256 => %x\n", sha512.Sum512_256(data))
	fmt.Print("===\n")
	fmt.Printf("SHA3-224 => %x\n", sha3.Sum224(data))
	fmt.Print("===\n")
	fmt.Printf("SHA3-256 => %x\n", sha3.Sum256(data))
	fmt.Print("===\n")
	fmt.Printf("SHA3-384 => %x\n", sha3.Sum384(data))
	fmt.Print("===\n")
	fmt.Printf("SHA3-512 => %x\n", sha3.Sum512(data))
}
