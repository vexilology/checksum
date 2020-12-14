package src

import (
  "fmt"
  "log"
  "strconv"
  "hash/fnv"
  "io/ioutil"
  "hash/crc32"
  "hash/crc64"
  "hash/adler32"
  "crypto/md5"
  "crypto/sha1"
  "crypto/sha256"
  "crypto/sha512"
  "encoding/base32"
  "encoding/base64"

  "golang.org/x/crypto/md4"
  "golang.org/x/crypto/blake2s"
  "golang.org/x/crypto/blake2b"
  "golang.org/x/crypto/ripemd160"
  "golang.org/x/crypto/sha3"

  "github.com/cxmcc/tiger"
  "github.com/ebfe/keccak"
  "github.com/htruong/go-md2"
)

func CheckFile(foundFile string) string {
  fl, err := ioutil.ReadFile(foundFile)
  if err != nil {
    log.Fatal(err)
  }

  newshake128 := keccak.NewSHAKE128(32)
  newshake128.Write([]byte(fl))
  newshake256 := keccak.NewSHAKE256(64)
  newshake256.Write([]byte(fl))

  newkeccak224 := keccak.New224()
  newkeccak224.Write([]byte(fl))
  newkeccak384 := keccak.New384()
  newkeccak384.Write([]byte(fl))
  newkeccak256 := sha3.NewLegacyKeccak256()
  newkeccak256.Write([]byte(fl))
  newkeccak512 := sha3.NewLegacyKeccak512()
  newkeccak512.Write([]byte(fl))

  firstcrc32 := crc32.ChecksumIEEE([]byte(fl))
  secondcrc32 := strconv.FormatUint(uint64(firstcrc32), 16)
  crc64ecma := crc64.New(crc64.MakeTable(crc64.ECMA))
  crc64ecma.Write([]byte(fl))
  crc64iso := crc64.New(crc64.MakeTable(crc64.ISO))
  crc64iso.Write([]byte(fl))

  firstadler32 := adler32.Checksum([]byte(fl))
  secondadler32 := strconv.FormatUint(uint64(firstadler32), 16)

  xbase32 := base32.StdEncoding.EncodeToString(fl)
  xbase64 := base64.StdEncoding.EncodeToString(fl)

  xMD2 := md2.New()
  xMD2.Write([]byte(fl))

  xMD4 := md4.New()
  xMD4.Write([]byte(fl))

  xRIPEMD160 := ripemd160.New()
  xRIPEMD160.Write([]byte(fl))

  xTiger := tiger.New()
  xTiger.Write([]byte(fl))

  xFnv32 := fnv.New32()
  xFnv32.Write([]byte(fl))
  xFnv64 := fnv.New64()
  xFnv64.Write([]byte(fl))
  xFnv128 := fnv.New128()
  xFnv128.Write([]byte(fl))
  xFnv32a := fnv.New32a()
  xFnv32a.Write([]byte(fl))
  xFnv64a := fnv.New64a()
  xFnv64a.Write([]byte(fl))
  xFnv128a := fnv.New128a()
  xFnv128a.Write([]byte(fl))

  fmt.Println("-----------------------")
  fmt.Printf("Length :: %d bytes\n", len(fl))
  fmt.Println("-----------------------")
  fmt.Println("BASE32 ::", xbase32)
  fmt.Println("BASE64 ::", xbase64)
  fmt.Println("-----------------------")
  fmt.Printf("TIGER192,3 :: %x\n", xTiger.Sum(nil))
  fmt.Printf("SHAKE128-256 :: %x\n", newshake128.Sum(nil))
  fmt.Printf("SHAKE256-512 :: %x\n", newshake256.Sum(nil))
  fmt.Printf("KECCAK224 :: %x\n", newkeccak224.Sum(nil))
  fmt.Printf("KECCAK256 :: %x\n", newkeccak256.Sum(nil))
  fmt.Printf("KECCAK384 :: %x\n", newkeccak384.Sum(nil))
  fmt.Printf("KECCAK512 :: %x\n", newkeccak512.Sum(nil))
  fmt.Println("CRC32-ieee ::", secondcrc32)
  fmt.Printf("CRC64-ecma :: %x\n", crc64ecma.Sum(nil))
  fmt.Printf("CRC64-iso  :: %x\n", crc64iso.Sum(nil))
  fmt.Println("ADLER32 ::", secondadler32)
  fmt.Printf("BLAKE2S-256 :: %x\n", blake2s.Sum256(fl))
  fmt.Printf("BLAKE2B-256 :: %x\n", blake2b.Sum256(fl))
  fmt.Printf("BLAKE2B-384 :: %x\n", blake2b.Sum384(fl))
  fmt.Printf("BLAKE2B-512 :: %x\n", blake2b.Sum512(fl))
  fmt.Printf("RIPEMD160 :: %x\n", xRIPEMD160.Sum(nil))
  fmt.Printf("MD2 :: %x\n", xMD2.Sum([]byte(nil)))
  fmt.Printf("MD4 :: %x\n", xMD4.Sum([]byte(nil)))
  fmt.Printf("MD5 :: %x\n", md5.Sum(fl))
  fmt.Printf("SHA1   :: %x\n", sha1.Sum(fl))
  fmt.Printf("SHA224 :: %x\n", sha256.Sum224(fl))
  fmt.Printf("SHA256 :: %x\n", sha256.Sum256(fl))
  fmt.Printf("SHA384 :: %x\n", sha512.Sum384(fl))
  fmt.Printf("SHA512 :: %x\n", sha512.Sum512(fl))
  fmt.Printf("SHA512_224 :: %x\n", sha512.Sum512_224(fl))
  fmt.Printf("SHA512_256 :: %x\n", sha512.Sum512_256(fl))
  fmt.Printf("SHA3-224 :: %x\n", sha3.Sum224(fl))
  fmt.Printf("SHA3-256 :: %x\n", sha3.Sum256(fl))
  fmt.Printf("SHA3-384 :: %x\n", sha3.Sum384(fl))
  fmt.Printf("SHA3-512 :: %x\n", sha3.Sum512(fl))
  fmt.Printf("FNV1-32  :: %x\n", xFnv32.Sum(nil))
  fmt.Printf("FNV1-64  :: %x\n", xFnv64.Sum(nil))
  fmt.Printf("FNV1-128 :: %x\n", xFnv128.Sum(nil))
  fmt.Printf("FNV1a-32  :: %x\n", xFnv32a.Sum(nil))
  fmt.Printf("FNV1a-64  :: %x\n", xFnv64a.Sum(nil))
  fmt.Printf("FNV1a-128 :: %x\n", xFnv128a.Sum(nil))
  fmt.Println("-----------------------")
  fmt.Println("OK")
  return foundFile
}
