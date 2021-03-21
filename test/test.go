package test

import (
  "fmt"
  "crypto/sha256"

  "golang.org/x/crypto/md4"
  "golang.org/x/crypto/blake2s"
  "golang.org/x/crypto/blake2b"
  "golang.org/x/crypto/ripemd160"
  "golang.org/x/crypto/sha3"

  "github.com/cxmcc/tiger"
  "github.com/htruong/go-md2"
)

func FoundSHA3_256(input_string string) string {
  return fmt.Sprintf("%x", sha3.Sum256([]byte(input_string)))
}

func FoundTIGER192(input_string string) string {
  f_tiger192 := tiger.New()
  f_tiger192.Write([]byte(input_string))
  return fmt.Sprintf("%x", f_tiger192.Sum(nil))
}

func FoundRIPEMD160(input_string string) string {
  f_ripemd := ripemd160.New()
  f_ripemd.Write([]byte(input_string))
  return fmt.Sprintf("%x", f_ripemd.Sum(nil))
}

func FoundBLAKE2s256(input_string string) string {
  return fmt.Sprintf("%x", blake2s.Sum256([]byte(input_string)))
}

func FoundBLAKE2b256(input_string string) string {
  return fmt.Sprintf("%x", blake2b.Sum256([]byte(input_string)))
}

func FoundKECCAK256(input_string string) string {
  newkeccak256 := sha3.NewLegacyKeccak256()
  newkeccak256.Write([]byte(input_string))
  return fmt.Sprintf("%x", newkeccak256.Sum(nil))
}

func FoundMD2(input_string string) string {
  f_MD2 := md2.New()
  f_MD2.Write([]byte(input_string))
  return fmt.Sprintf("%x", f_MD2.Sum(nil))
}

func FoundMD4(input_string string) string {
  f_MD4 := md4.New()
  f_MD4.Write([]byte(input_string))
  return fmt.Sprintf("%x", f_MD4.Sum([]byte(nil)))
}

func FoundSHA256(input_string string) string {
  return fmt.Sprintf("%x", sha256.Sum256([]byte(input_string)))
}
