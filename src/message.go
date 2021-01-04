package src

import (
  "io"
  "strings"
  "log"
  "fmt"
  "strconv"
  "hash/fnv"
  "hash/crc32"
  "hash/crc64"
  "hash/adler32"
  "crypto/md5"
  "crypto/sha1"
  "crypto/sha256"
  "crypto/sha512"
  "encoding/ascii85"
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

func CheckMessage(foundString string) string {
  input_string := foundString
  msg := make([]byte, len(foundString))
  if _, err := io.ReadFull(strings.NewReader(input_string), msg); err != nil {
    log.Fatal(err)
  }

  firstShake128 := keccak.NewSHAKE128(32)
  firstShake128.Write([]byte(msg))
  secondShake128 := fmt.Sprintf("%x", firstShake128.Sum(nil))
  firstShake256 := keccak.NewSHAKE256(64)
  firstShake256.Write([]byte(msg))
  secondShake256 := fmt.Sprintf("%x", firstShake256.Sum(nil))

  newkeccak224 := keccak.New224()
  newkeccak224.Write([]byte(msg))
  lastkeccak224 := fmt.Sprintf("%x", newkeccak224.Sum(nil))
  newkeccak384 := keccak.New384()
  newkeccak384.Write([]byte(msg))
  lastkeccak384 := fmt.Sprintf("%x", newkeccak384.Sum(nil))
  newkeccak256 := sha3.NewLegacyKeccak256()
  newkeccak256.Write([]byte(msg))
  lastkeccak256 := fmt.Sprintf("%x", newkeccak256.Sum(nil))
  newkeccak512 := sha3.NewLegacyKeccak512()
  newkeccak512.Write([]byte(msg))
  lastkeccak512 := fmt.Sprintf("%x", newkeccak512.Sum(nil))

  firstCRC32 := crc32.ChecksumIEEE([]byte(msg))
  secondCRC32 := strconv.FormatUint(uint64(firstCRC32), 16)
  firstCRC64ecma := crc64.New(crc64.MakeTable(crc64.ECMA))
  firstCRC64ecma.Write([]byte(msg))
  secondCRC64ecma := fmt.Sprintf("%x", firstCRC64ecma.Sum(nil))
  firstCRC64iso := crc64.New(crc64.MakeTable(crc64.ISO))
  firstCRC64iso.Write([]byte(msg))
  secondCRC64iso := fmt.Sprintf("%x", firstCRC64iso.Sum(nil))

  firstADLER32 := adler32.Checksum([]byte(msg))
  secondADLER32 := strconv.FormatUint(uint64(firstADLER32), 16)

  firstmd2 := md2.New()
  firstmd2.Write([]byte(msg))
  secondmd2 := fmt.Sprintf("%x", firstmd2.Sum([]byte(nil)))

  firstMD4 := md4.New()
  firstMD4.Write([]byte(msg))
  secondMD4 := fmt.Sprintf("%x", firstMD4.Sum([]byte(nil)))

  MD5 := fmt.Sprintf("%x", md5.Sum([]byte(msg)))

  SHA1 := fmt.Sprintf("%x", sha1.Sum([]byte(msg)))

  SHA224 := fmt.Sprintf("%x", sha256.Sum224([]byte(msg)))
  SHA256 := fmt.Sprintf("%x", sha256.Sum256([]byte(msg)))
  SHA384 := fmt.Sprintf("%x", sha512.Sum384([]byte(msg)))
  SHA512 := fmt.Sprintf("%x", sha512.Sum512([]byte(msg)))
  SHA512_224 := fmt.Sprintf("%x", sha512.Sum512_224([]byte(msg)))
  SHA512_256 := fmt.Sprintf("%x", sha512.Sum512_256([]byte(msg)))

  ASCII85_first := make([]byte, 10000, 10000)
  ascii85.Encode(ASCII85_first, []byte(msg))
  ASCII85 := string(ASCII85_first)

  BASE32 := base32.StdEncoding.EncodeToString([]byte(msg))
  BASE64 := base64.StdEncoding.EncodeToString([]byte(msg))

  BLAKE2S256 := fmt.Sprintf("%x", blake2s.Sum256([]byte(msg)))
  BLAKE2B256 := fmt.Sprintf("%x", blake2b.Sum256([]byte(msg)))
  BLAKE2B384 := fmt.Sprintf("%x", blake2b.Sum384([]byte(msg)))
  BLAKE2B512 := fmt.Sprintf("%x", blake2b.Sum512([]byte(msg)))

  firstRIPEMD160 := ripemd160.New()
  firstRIPEMD160.Write([]byte(msg))
  secondRIPEMD160 := fmt.Sprintf("%x", firstRIPEMD160.Sum(nil))

  SHA3_224 := fmt.Sprintf("%x", sha3.Sum224([]byte(msg)))
  SHA3_256 := fmt.Sprintf("%x", sha3.Sum256([]byte(msg)))
  SHA3_384 := fmt.Sprintf("%x", sha3.Sum384([]byte(msg)))
  SHA3_512 := fmt.Sprintf("%x", sha3.Sum512([]byte(msg)))

  firstTiger := tiger.New()
  firstTiger.Write([]byte(msg))
  secondTiger := fmt.Sprintf("%x", firstTiger.Sum(nil))

  fnv32 := fnv.New32()
  fnv32.Write([]byte(msg))
  fnv32last := fmt.Sprintf("%x", fnv32.Sum(nil))
  fnv64 := fnv.New64()
  fnv64.Write([]byte(msg))
  fnv64last := fmt.Sprintf("%x", fnv64.Sum(nil))
  fnv128 := fnv.New128()
  fnv128.Write([]byte(msg))
  fnv128last := fmt.Sprintf("%x", fnv128.Sum(nil))
  fnv32a := fnv.New32a()
  fnv32a.Write([]byte(msg))
  fnv32alast := fmt.Sprintf("%x", fnv32a.Sum(nil))
  fnv64a := fnv.New64a()
  fnv64a.Write([]byte(msg))
  fnv64alast := fmt.Sprintf("%x", fnv64a.Sum(nil))
  fnv128a := fnv.New128a()
  fnv128a.Write([]byte(msg))
  fnv128alast := fmt.Sprintf("%x", fnv128a.Sum(nil))

  fmt.Println("-----------------------")
  fmt.Println("ASCII85 ::", ASCII85)
  fmt.Println("BASE32  ::", BASE32)
  fmt.Println("BASE64  ::", BASE64)
  fmt.Println("-----------------------")
  fmt.Println("TIGER192,3 ::", secondTiger)
  fmt.Println("SHAKE128-256 ::", secondShake128)
  fmt.Println("SHAKE256-512 ::", secondShake256)
  fmt.Println("KECCAK224 ::", lastkeccak224)
  fmt.Println("KECCAK256 ::", lastkeccak256)
  fmt.Println("KECCAK384 ::", lastkeccak384)
  fmt.Println("KECCAK512 ::", lastkeccak512)
  fmt.Println("CRC32-ieee ::", secondCRC32)
  fmt.Println("CRC64-ecma ::", secondCRC64ecma)
  fmt.Println("CRC64-iso  ::", secondCRC64iso)
  fmt.Println("ADLER32 ::", secondADLER32)
  fmt.Println("BLAKE2S-256 ::", BLAKE2S256)
  fmt.Println("BLAKE2B-256 ::", BLAKE2B256)
  fmt.Println("BLAKE2B-384 ::", BLAKE2B384)
  fmt.Println("BLAKE2B-512 ::", BLAKE2B512)
  fmt.Println("RIPEMD160 ::", secondRIPEMD160)
  fmt.Println("MD2 ::", secondmd2)
  fmt.Println("MD4 ::", secondMD4)
  fmt.Println("MD5 ::", MD5)
  fmt.Println("SHA1   ::", SHA1)
  fmt.Println("SHA224 ::", SHA224)
  fmt.Println("SHA256 ::", SHA256)
  fmt.Println("SHA384 ::", SHA384)
  fmt.Println("SHA512 ::", SHA512)
  fmt.Println("SHA512_224 ::", SHA512_224)
  fmt.Println("SHA512_256 ::", SHA512_256)
  fmt.Println("SHA3-224   ::", SHA3_224)
  fmt.Println("SHA3-256   ::", SHA3_256)
  fmt.Println("SHA3-384   ::", SHA3_384)
  fmt.Println("SHA3-512   ::", SHA3_512)
  fmt.Println("FNV1-32  ::", fnv32last)
  fmt.Println("FNV1-62  ::", fnv64last)
  fmt.Println("FNV1-128 ::", fnv128last)
  fmt.Println("FNV1a-32  ::", fnv32alast)
  fmt.Println("FNV1a-64  ::", fnv64alast)
  fmt.Println("FNV1a-128 ::", fnv128alast)
  fmt.Println("-----------------------")
  fmt.Println("OK")
  return foundString
}
