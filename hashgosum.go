package main

import (
  "fmt"
  "flag"
)

var (
  stringFound = flag.String("m", "", "found string")
  fileFound = flag.String("f", "", "found file")
  algorithmName = flag.String("a", "", "algorithm name")
  saveR = flag.Bool("s", false, "save result in file")
  deleteF = flag.Bool("d", false, "delete created file")
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
  } else if *deleteF {
    DeleteFile()
  } else if *saveR {
    SaveOutput()
  } else if *fileFound != "" {
    CasefullFile()
  } else if *stringFound != "" {
    CasefullMessage()
  } else {
    fmt.Println("Empty message, try again.")
  }
}

func main() {
  parseFlags()
}
