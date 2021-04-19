package main

import (
  "log"
  "fmt"
  "flag"
  "io/ioutil"

  "github.com/vexilology/hashgosum/algorithm"
)

var (
  isHelp        = flag.Bool("h", false, "help")
  stringFound   = flag.String("s", "", "found string")
  fileFound     = flag.String("f", "", "found file")
  algorithmName = flag.String("a", "", "algorithm name")
)

var algorithm_list = []string{
  "sha3_224", "sha3_256", "sha3_384", "sha3_512",
  "tiger192", "ripemd160", "adler32",
  "blake2s256", "blake2b256", "blake2b384", "blake2b512",
  "keccak224", "keccak256", "keccak384", "keccak512",
  "crc32ieee", "crc64ecma", "crc64iso",
  "shake128_224", "shake128_256", "shake128_384", "shake128_512",
  "shake256_224", "shake256_256", "shake256_384", "shake256_512",
  "md2", "md4", "md5",
  "fnv32", "fnv32a", "fnv64", "fnv64a", "fnv128", "fnv128a",
  "sha1", "sha224", "sha256", "sha384", "sha512",
  "sha512_224", "sha512_256",
  "base32", "base64",
}

func fileToString(foundFile string) string {
  f_file, err := ioutil.ReadFile(foundFile)
  if err != nil {
    log.Fatal("File not found.")
  }
  s_file := string(f_file)
  return s_file
}

// FIXME defaultError() now not used
func defaultError() {
  if *algorithmName == "" {
    fmt.Println("Unknown algorithm, try again.")
  }
}

func parseFlags() {
  flag.Parse()

  h := map[string]func(string) (int, error) {
    "md2": func(s string) (int, error) {
      return algorithm.FoundMD2(s)
    },
    "md4": func(s string) (int, error) {
      return algorithm.FoundMD4(s)
    },
    "md5": func(s string) (int, error) {
      return algorithm.FoundMD5(s)
    },
    "crc32ieee": func(s string) (int, error) {
      return algorithm.FoundCRC32ieee(s)
    },
    "crc64ecma": func(s string) (int, error) {
      return algorithm.FoundCRC64ecma(s)
    },
    "crc64iso": func(s string) (int, error) {
      return algorithm.FoundCRC64iso(s)
    },
    "adler32": func(s string) (int, error) {
      return algorithm.FoundADLER32(s)
    },
    "tiger192": func(s string) (int, error) {
      return algorithm.FoundTIGER192(s)
    },
    "ripemd160": func(s string) (int, error) {
      return algorithm.FoundRIPEMD160(s)
    },
    "keccak224": func(s string) (int, error) {
      return algorithm.FoundKECCAK224(s)
    },
    "keccak256": func(s string) (int, error) {
      return algorithm.FoundKECCAK256(s)
    },
    "keccak384": func(s string) (int, error) {
      return algorithm.FoundKECCAK384(s)
    },
    "keccak512": func(s string) (int, error) {
      return algorithm.FoundKECCAK512(s)
    },
    "blake2s256": func(s string) (int, error) {
      return algorithm.FoundBLAKE2s256(s)
    },
    "blake2b256": func(s string) (int, error) {
      return algorithm.FoundBLAKE2b256(s)
    },
    "blake2b384": func(s string) (int, error) {
      return algorithm.FoundBLAKE2b384(s)
    },
    "blake2b512": func(s string) (int, error) {
      return algorithm.FoundBLAKE2b512(s)
    },
    "sha1": func(s string) (int, error) {
      return algorithm.FoundSHA1(s)
    },
    "sha256": func(s string) (int, error) {
      return algorithm.FoundSHA256(s)
    },
    "sha224": func(s string) (int, error) {
      return algorithm.FoundSHA224(s)
    },
    "sha384": func(s string) (int, error) {
      return algorithm.FoundSHA384(s)
    },
    "sha512": func(s string) (int, error) {
      return algorithm.FoundSHA512(s)
    },
    "sha3_224": func(s string) (int, error) {
      return algorithm.FoundSHA3_224(s)
    },
    "sha3_256": func(s string) (int, error) {
      return algorithm.FoundSHA3_256(s)
    },
    "sha3_384": func(s string) (int, error) {
      return algorithm.FoundSHA3_384(s)
    },
    "sha3_512": func(s string) (int, error) {
      return algorithm.FoundSHA3_512(s)
    },
    "sha512_224": func(s string) (int, error) {
      return algorithm.FoundSHA512_224(s)
    },
    "sha512_256": func(s string) (int, error) {
      return algorithm.FoundSHA512_256(s)
    },
    "shake128_224": func(s string) (int, error) {
      return algorithm.FoundSHAKE128_224(s)
    },
    "shake128_256": func(s string) (int, error) {
      return algorithm.FoundSHAKE128_256(s)
    },
    "shake128_384": func(s string) (int, error) {
      return algorithm.FoundSHAKE128_384(s)
    },
    "shake128_512": func(s string) (int, error) {
      return algorithm.FoundSHAKE128_512(s)
    },
    "shake256_224": func(s string) (int, error) {
      return algorithm.FoundSHAKE256_224(s)
    },
    "shake256_256": func(s string) (int, error) {
      return algorithm.FoundSHAKE256_256(s)
    },
    "shake256_384": func(s string) (int, error) {
      return algorithm.FoundSHAKE256_384(s)
    },
    "shake256_512": func(s string) (int, error) {
      return algorithm.FoundSHAKE256_512(s)
    },
    "fnv32": func(s string) (int, error) {
      return algorithm.FoundFNV32(s)
    },
    "fnv32a": func(s string) (int, error) {
      return algorithm.FoundFNV32a(s)
    },
    "fnv64": func(s string) (int, error) {
      return algorithm.FoundFNV64(s)
    },
    "fnv64a": func(s string) (int, error) {
      return algorithm.FoundFNV64a(s)
    },
    "fnv128": func(s string) (int, error) {
      return algorithm.FoundFNV128(s)
    },
    "fnv128a": func(s string) (int, error) {
      return algorithm.FoundFNV128a(s)
    },
    "base32": func(s string) (int, error) {
      return algorithm.FoundBASE32(s)
    },
    "base64": func(s string) (int, error) {
      return algorithm.FoundBASE64(s)
    },
  }

  if *isHelp {
    fmt.Println("Available:", []string(algorithm_list))
    flag.PrintDefaults()
  } else if *fileFound != "" {
    // FIXME there are always 2 digits at the end. hash sum is correct
    resultF, err := (h[*algorithmName](fileToString(*fileFound)))
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(resultF)
  } else if *stringFound != "" {
    // FIXME there are always 2 digits at the end. hash sum is correct
    resultS, err := (h[*algorithmName](*stringFound))
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(resultS)
  } else {
    fmt.Println("Empty message, try again.")
  }
}

func main() {
  parseFlags()
}
