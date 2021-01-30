package main

import (
  "fmt"
  "flag"
  "./algorithm"
)

var (
  stringFound = flag.String("m", "", "found string")
  fileFound = flag.String("f", "", "found file")
  algorithmName = flag.String("a", "", "algorithm name")
  isHelp = flag.Bool("h", false, "help")
)

func parseFlags() {
  flag.Parse()

  if *isHelp {
    fmt.Println("Available:",
                 "sha3_224, sha3_256, sha3_384, sha3_512,",
                 "tiger192,",
                 "ripemd160,",
                 "blake2s256,",
                 "blake2b256, blake2b384, blake2b512,",
                 "keccak224, keccak256, keccak384, keccak512,",
                 "adler32,",
                 "crc32ieee,",
                 "crc64ecma, crc64iso,",
                 "shake128_224, shake128_256, shake128_384, shake128_512,",
                 "shake256_224, shake256_256, shake256_384, shake256_512,",
                 "md2, md4, md5,",
                 "fnv32, fnv32a,",
                 "fnv64, fnv64a,",
                 "fnv128, fnv128a,",
                 "sha1, sha224, sha256, sha384, sha512,",
                 "sha512_224, sha512_256,",
                 "base32, base64.",
               )
    flag.PrintDefaults()
  } else if *fileFound != "" {
    switch *algorithmName {
    case "sha3_224":
      algorithm.FoundSHA3_224(CheckFileSolo(*fileFound))
    case "sha3_256":
      algorithm.FoundSHA3_256(CheckFileSolo(*fileFound))
    case "sha3_384":
      algorithm.FoundSHA3_384(CheckFileSolo(*fileFound))
    case "sha3_512":
      algorithm.FoundSHA3_512(CheckFileSolo(*fileFound))
    case "tiger192":
      algorithm.FoundTIGER192(CheckFileSolo(*fileFound))
    case "ripemd160":
      algorithm.FoundRIPEMD160(CheckFileSolo(*fileFound))
    case "blake2s256":
      algorithm.FoundBLAKE2s256(CheckFileSolo(*fileFound))
    case "blake2b256":
      algorithm.FoundBLAKE2b256(CheckFileSolo(*fileFound))
    case "blake2b384":
      algorithm.FoundBLAKE2b384(CheckFileSolo(*fileFound))
    case "blake2b512":
      algorithm.FoundBLAKE2b512(CheckFileSolo(*fileFound))
    case "keccak224":
      algorithm.FoundKECCAK224(CheckFileSolo(*fileFound))
    case "keccak256":
      algorithm.FoundKECCAK256(CheckFileSolo(*fileFound))
    case "keccak384":
      algorithm.FoundKECCAK384(CheckFileSolo(*fileFound))
    case "keccak512":
      algorithm.FoundKECCAK512(CheckFileSolo(*fileFound))
    case "adler32":
      algorithm.FoundADLER32(CheckFileSolo(*fileFound))
    case "crc32ieee":
      algorithm.FoundCRC32ieee(CheckFileSolo(*fileFound))
    case "crc64ecma":
      algorithm.FoundCRC64ecma(CheckFileSolo(*fileFound))
    case "crc64iso":
      algorithm.FoundCRC64iso(CheckFileSolo(*fileFound))
    case "shake128_224":
      algorithm.FoundSHAKE128_224(CheckFileSolo(*fileFound))
    case "shake128_256":
      algorithm.FoundSHAKE128_256(CheckFileSolo(*fileFound))
    case "shake128_384":
      algorithm.FoundSHAKE128_384(CheckFileSolo(*fileFound))
    case "shake128_512":
      algorithm.FoundSHAKE128_512(CheckFileSolo(*fileFound))
    case "shake256_224":
      algorithm.FoundSHAKE256_224(CheckFileSolo(*fileFound))
    case "shake256_256":
      algorithm.FoundSHAKE256_256(CheckFileSolo(*fileFound))
    case "shake256_384":
      algorithm.FoundSHAKE256_384(CheckFileSolo(*fileFound))
    case "shake256_512":
      algorithm.FoundSHAKE256_512(CheckFileSolo(*fileFound))
    case "md2":
      algorithm.FoundMD2(CheckFileSolo(*fileFound))
    case "md4":
      algorithm.FoundMD4(CheckFileSolo(*fileFound))
    case "md5":
      algorithm.FoundMD5(CheckFileSolo(*fileFound))
    case "fnv32":
      algorithm.FoundFNV32(CheckFileSolo(*fileFound))
    case "fnv32a":
      algorithm.FoundFNV32a(CheckFileSolo(*fileFound))
    case "fnv64":
      algorithm.FoundFNV64(CheckFileSolo(*fileFound))
    case "fnv64a":
      algorithm.FoundFNV64a(CheckFileSolo(*fileFound))
    case "fnv128":
      algorithm.FoundFNV128(CheckFileSolo(*fileFound))
    case "fnv128a":
      algorithm.FoundFNV128a(CheckFileSolo(*fileFound))
    case "sha1":
      algorithm.FoundSHA1(CheckFileSolo(*fileFound))
    case "sha224":
      algorithm.FoundSHA224(CheckFileSolo(*fileFound))
    case "sha256":
      algorithm.FoundSHA256(CheckFileSolo(*fileFound))
    case "sha384":
      algorithm.FoundSHA384(CheckFileSolo(*fileFound))
    case "sha512":
      algorithm.FoundSHA512(CheckFileSolo(*fileFound))
    case "sha512_224":
      algorithm.FoundSHA512_224(CheckFileSolo(*fileFound))
    case "sha512_256":
      algorithm.FoundSHA512_256(CheckFileSolo(*fileFound))
    case "base32":
      algorithm.FoundBASE32(CheckFileSolo(*fileFound))
    case "base64":
      algorithm.FoundBASE64(CheckFileSolo(*fileFound))
    default:
      CheckFileFull(*fileFound)
    }
  } else if *stringFound != "" {
    switch *algorithmName {
    case "sha3_224":
      algorithm.FoundSHA3_224(*stringFound)
    case "sha3_256":
      algorithm.FoundSHA3_256(*stringFound)
    case "sha3_384":
      algorithm.FoundSHA3_384(*stringFound)
    case "sha3_512":
      algorithm.FoundSHA3_512(*stringFound)
    case "tiger192":
      algorithm.FoundTIGER192(*stringFound)
    case "ripemd160":
      algorithm.FoundRIPEMD160(*stringFound)
    case "blake2s256":
      algorithm.FoundBLAKE2s256(*stringFound)
    case "blake2b256":
      algorithm.FoundBLAKE2b256(*stringFound)
    case "blake2b384":
      algorithm.FoundBLAKE2b384(*stringFound)
    case "blake2b512":
      algorithm.FoundBLAKE2b512(*stringFound)
    case "keccak224":
      algorithm.FoundKECCAK224(*stringFound)
    case "keccak256":
      algorithm.FoundKECCAK256(*stringFound)
    case "keccak384":
      algorithm.FoundKECCAK384(*stringFound)
    case "keccak512":
      algorithm.FoundKECCAK512(*stringFound)
    case "adler32":
      algorithm.FoundADLER32(*stringFound)
    case "crc32ieee":
      algorithm.FoundCRC32ieee(*stringFound)
    case "crc64ecma":
      algorithm.FoundCRC64ecma(*stringFound)
    case "crc64iso":
      algorithm.FoundCRC64iso(*stringFound)
    case "shake128_224":
      algorithm.FoundSHAKE128_224(*stringFound)
    case "shake128_256":
      algorithm.FoundSHAKE128_256(*stringFound)
    case "shake128_384":
      algorithm.FoundSHAKE128_384(*stringFound)
    case "shake128_512":
      algorithm.FoundSHAKE128_512(*stringFound)
    case "shake256_224":
      algorithm.FoundSHAKE256_224(*stringFound)
    case "shake256_256":
      algorithm.FoundSHAKE256_256(*stringFound)
    case "shake256_384":
      algorithm.FoundSHAKE256_384(*stringFound)
    case "shake256_512":
      algorithm.FoundSHAKE256_512(*stringFound)
    case "md2":
      algorithm.FoundMD2(*stringFound)
    case "md4":
      algorithm.FoundMD4(*stringFound)
    case "md5":
      algorithm.FoundMD5(*stringFound)
    case "fnv32":
      algorithm.FoundFNV32(*stringFound)
    case "fnv32a":
      algorithm.FoundFNV32a(*stringFound)
    case "fnv64":
      algorithm.FoundFNV64(*stringFound)
    case "fnv64a":
      algorithm.FoundFNV64a(*stringFound)
    case "fnv128":
      algorithm.FoundFNV128(*stringFound)
    case "fnv128a":
      algorithm.FoundFNV128a(*stringFound)
    case "sha1":
      algorithm.FoundSHA1(*stringFound)
    case "sha224":
      algorithm.FoundSHA224(*stringFound)
    case "sha256":
      algorithm.FoundSHA256(*stringFound)
    case "sha384":
      algorithm.FoundSHA384(*stringFound)
    case "sha512":
      algorithm.FoundSHA512(*stringFound)
    case "sha512_224":
      algorithm.FoundSHA512_224(*stringFound)
    case "sha512_256":
      algorithm.FoundSHA512_256(*stringFound)
    case "base32":
      algorithm.FoundBASE32(*stringFound)
    case "base64":
      algorithm.FoundBASE64(*stringFound)
    default:
      CheckMessageFull(*stringFound)
    }
  } else {
    fmt.Println("Empty message, try again.")
  }
}

func main() {
  parseFlags()
}
