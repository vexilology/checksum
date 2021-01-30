package main

import (
  "io/ioutil"
  "./algorithm"
)

func CheckFileSolo(foundFile string) string {
  f_file, err := ioutil.ReadFile(foundFile)
  if err != nil {
    panic(err)
  }

  s_file := string(f_file)
  return s_file
}

func CheckFileFull(foundFile string) string {
  f_file, err := ioutil.ReadFile(foundFile)
  if err != nil {
    panic(err)
  }

  s_file := string(f_file)

  algorithm.FoundBASE32(s_file)
  algorithm.FoundBASE64(s_file)
  algorithm.FoundTIGER192(s_file)
  algorithm.FoundSHAKE128_224(s_file)
  algorithm.FoundSHAKE128_256(s_file)
  algorithm.FoundSHAKE128_384(s_file)
  algorithm.FoundSHAKE128_512(s_file)
  algorithm.FoundSHAKE256_224(s_file)
  algorithm.FoundSHAKE256_256(s_file)
  algorithm.FoundSHAKE256_384(s_file)
  algorithm.FoundSHAKE256_512(s_file)
  algorithm.FoundKECCAK224(s_file)
  algorithm.FoundKECCAK256(s_file)
  algorithm.FoundKECCAK384(s_file)
  algorithm.FoundKECCAK512(s_file)
  algorithm.FoundCRC32ieee(s_file)
  algorithm.FoundCRC64ecma(s_file)
  algorithm.FoundCRC64iso(s_file)
  algorithm.FoundADLER32(s_file)
  algorithm.FoundBLAKE2s256(s_file)
  algorithm.FoundBLAKE2b256(s_file)
  algorithm.FoundBLAKE2b384(s_file)
  algorithm.FoundBLAKE2b512(s_file)
  algorithm.FoundRIPEMD160(s_file)
  algorithm.FoundMD2(s_file)
  algorithm.FoundMD4(s_file)
  algorithm.FoundMD5(s_file)
  algorithm.FoundSHA1(s_file)
  algorithm.FoundSHA224(s_file)
  algorithm.FoundSHA256(s_file)
  algorithm.FoundSHA384(s_file)
  algorithm.FoundSHA512(s_file)
  algorithm.FoundSHA512_224(s_file)
  algorithm.FoundSHA512_256(s_file)
  algorithm.FoundSHA3_224(s_file)
  algorithm.FoundSHA3_256(s_file)
  algorithm.FoundSHA3_384(s_file)
  algorithm.FoundSHA3_512(s_file)
  algorithm.FoundFNV32(s_file)
  algorithm.FoundFNV32a(s_file)
  algorithm.FoundFNV64(s_file)
  algorithm.FoundFNV64a(s_file)
  algorithm.FoundFNV128(s_file)
  algorithm.FoundFNV128a(s_file)
  return foundFile
}
