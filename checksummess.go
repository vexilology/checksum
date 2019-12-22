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
	"golang.org/x/crypto/sha3"
	"github.com/cxmcc/tiger"
)

func main() {

	var data string

	fmt.Print("Enter your message: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		data = scanner.Text()
	}

	newkeccak256 := sha3.NewLegacyKeccak256()
	newkeccak256.Write([]byte(data))
	lastkeccak256 := fmt.Sprintf("%x", newkeccak256.Sum(nil))
	newkeccak512 := sha3.NewLegacyKeccak512()
	newkeccak512.Write([]byte(data))
	lastkeccak512 := fmt.Sprintf("%x", newkeccak512.Sum(nil))

	firstCRC32 := crc32.ChecksumIEEE([]byte(data))
	secondCRC32 := strconv.FormatUint(uint64(firstCRC32), 16)

	firstADLER32 := adler32.Checksum([]byte(data))
	secondADLER32 := strconv.FormatUint(uint64(firstADLER32), 16)

	firstMD4 := md4.New()
	firstMD4.Write([]byte(data))
	secondMD4 := fmt.Sprintf("%x", firstMD4.Sum([]byte(nil)))
	MD5 := fmt.Sprintf("%x", md5.Sum([]byte(data)))

	SHA1 := fmt.Sprintf("%x", sha1.Sum([]byte(data)))
	SHA224 := fmt.Sprintf("%x", sha256.Sum224([]byte(data)))
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

	SHA3_224 := fmt.Sprintf("%x", sha3.Sum224([]byte(data)))
	SHA3_256 := fmt.Sprintf("%x", sha3.Sum256([]byte(data)))
	SHA3_384 := fmt.Sprintf("%x", sha3.Sum384([]byte(data)))
	SHA3_512 := fmt.Sprintf("%x", sha3.Sum512([]byte(data)))

	firstTiger := tiger.New()
	firstTiger.Write([]byte(data))
	secondTiger := fmt.Sprintf("%x", firstTiger.Sum(nil))

	fmt.Print("-----------------------\n")
	fmt.Println("BASE32:", BASE32)
	fmt.Print("===\n")
	fmt.Println("BASE64:", BASE64)
	fmt.Print("===\n")
	fmt.Println("tiger192[rounds-3]:", secondTiger)
	fmt.Print("===\n")
	fmt.Println("KECCAK256:", lastkeccak256)
	fmt.Print("===\n")
	fmt.Println("KECCAK512:", lastkeccak512)
	fmt.Print("===\n")
	fmt.Println("CRC32:", secondCRC32)
	fmt.Print("===\n")
	fmt.Println("ADLER32:", secondADLER32)
	fmt.Print("===\n")
	fmt.Println("BLAKE2B-256:", BLAKE2B256)
	fmt.Print("===\n")
	fmt.Println("BLAKE2B-512:", BLAKE2B512)
	fmt.Print("===\n")
	fmt.Println("RIPEMD160:", secondRIPEMD160)
	fmt.Print("===\n")
	fmt.Println("MD4:", secondMD4)
	fmt.Print("===\n")
	fmt.Println("MD5:", MD5)
	fmt.Print("===\n")
	fmt.Println("SHA1:", SHA1)
	fmt.Print("===\n")
	fmt.Println("SHA224:", SHA224)
	fmt.Print("===\n")
	fmt.Println("SHA256:", SHA256)
	fmt.Print("===\n")
	fmt.Println("SHA384:", SHA384)
	fmt.Print("===\n")
	fmt.Println("SHA512:", SHA512)
	fmt.Print("===\n")
	fmt.Println("SHA3-224:", SHA3_224)
	fmt.Print("===\n")
	fmt.Println("SHA3-256:", SHA3_256)
	fmt.Print("===\n")
	fmt.Println("SHA3-384:", SHA3_384)
	fmt.Print("===\n")
	fmt.Println("SHA3-512:", SHA3_512)
}
