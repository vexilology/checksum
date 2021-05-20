package main

import "github.com/vexilology/hashgosum/algorithm"

var h = map[string]func(string) string {
  "md2": func(s string) string {
    return algorithm.FoundMD2(s)
  },
  "md4": func(s string) string {
    return algorithm.FoundMD4(s)
  },
  "md5": func(s string) string {
    return algorithm.FoundMD5(s)
  },
  "crc32ieee": func(s string) string {
    return algorithm.FoundCRC32ieee(s)
  },
  "crc64ecma": func(s string) string {
    return algorithm.FoundCRC64ecma(s)
  },
  "crc64iso": func(s string) string {
    return algorithm.FoundCRC64iso(s)
  },
  "adler32": func(s string) string {
    return algorithm.FoundADLER32(s)
  },
  "tiger192": func(s string) string {
    return algorithm.FoundTIGER192(s)
  },
  "ripemd160": func(s string) string {
    return algorithm.FoundRIPEMD160(s)
  },
  "keccak224": func(s string) string {
    return algorithm.FoundKECCAK224(s)
  },
  "keccak256": func(s string) string {
    return algorithm.FoundKECCAK256(s)
  },
  "keccak384": func(s string) string {
    return algorithm.FoundKECCAK384(s)
  },
  "keccak512": func(s string) string {
    return algorithm.FoundKECCAK512(s)
  },
  "blake2s256": func(s string) string {
    return algorithm.FoundBLAKE2s256(s)
  },
  "blake2b256": func(s string) string {
    return algorithm.FoundBLAKE2b256(s)
  },
  "blake2b384": func(s string) string {
    return algorithm.FoundBLAKE2b384(s)
  },
  "blake2b512": func(s string) string {
    return algorithm.FoundBLAKE2b512(s)
  },
  "sha1": func(s string) string {
    return algorithm.FoundSHA1(s)
  },
  "sha256": func(s string) string {
    return algorithm.FoundSHA256(s)
  },
  "sha224": func(s string) string {
    return algorithm.FoundSHA224(s)
  },
  "sha384": func(s string) string {
    return algorithm.FoundSHA384(s)
  },
  "sha512": func(s string) string {
    return algorithm.FoundSHA512(s)
  },
  "sha3_224": func(s string) string {
    return algorithm.FoundSHA3_224(s)
  },
  "sha3_256": func(s string) string {
    return algorithm.FoundSHA3_256(s)
  },
  "sha3_384": func(s string) string {
    return algorithm.FoundSHA3_384(s)
  },
  "sha3_512": func(s string) string {
    return algorithm.FoundSHA3_512(s)
  },
  "sha512_224": func(s string) string {
    return algorithm.FoundSHA512_224(s)
  },
  "sha512_256": func(s string) string {
    return algorithm.FoundSHA512_256(s)
  },
  "shake128_224": func(s string) string {
    return algorithm.FoundSHAKE128_224(s)
  },
  "shake128_256": func(s string) string {
    return algorithm.FoundSHAKE128_256(s)
  },
  "shake128_384": func(s string) string {
    return algorithm.FoundSHAKE128_384(s)
  },
  "shake128_512": func(s string) string {
    return algorithm.FoundSHAKE128_512(s)
  },
  "shake256_224": func(s string) string {
    return algorithm.FoundSHAKE256_224(s)
  },
  "shake256_256": func(s string) string {
    return algorithm.FoundSHAKE256_256(s)
  },
  "shake256_384": func(s string) string {
    return algorithm.FoundSHAKE256_384(s)
  },
  "shake256_512": func(s string) string {
    return algorithm.FoundSHAKE256_512(s)
  },
  "fnv32": func(s string) string {
    return algorithm.FoundFNV32(s)
  },
  "fnv32a": func(s string) string {
    return algorithm.FoundFNV32a(s)
  },
  "fnv64": func(s string) string {
    return algorithm.FoundFNV64(s)
  },
  "fnv64a": func(s string) string {
    return algorithm.FoundFNV64a(s)
  },
  "fnv128": func(s string) string {
    return algorithm.FoundFNV128(s)
  },
  "fnv128a": func(s string) string {
    return algorithm.FoundFNV128a(s)
  },
  "base32": func(s string) string {
    return algorithm.FoundBASE32(s)
  },
  "base64": func(s string) string {
    return algorithm.FoundBASE64(s)
  },
}
