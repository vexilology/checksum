package main

import (
  "testing"

  "github.com/vexilology/hashgosum/test"
)

var (
  random_message string = "helloworld"
  sha256sum      string = "936a185caaa266bb9cbe981e9e05cb78cd732b0b3280eb944412bb6f8f8f07af"
  keccak256sum   string = "fa26db7ca85ead399216e7c6316bc50ed24393c3122b582735e7f3b0f91b93f0"
  md2sum         string = "0ae06562456c6e0e9736f4feda3a477b"
  tiger192sum    string = "5c9943da4a68b90dd5489ddef8d9eb6d71c9e0b00e2c4347"
  md4sum         string = "793033db97268fc9ceebde269797e54b"
  sha3_256sum    string = "92dad9443e4dd6d70a7f11872101ebff87e21798e4fbb26fa4bf590eb440e71b"
  ripemd160sum   string = "8a73c5438c28e79e696144fa869886f240cfaddb"
  blake2s256sum  string = "909052ce75dc79b301b0347ac904bae3b3117ce5abf963f16ee4c8ce08a3d407"
  blake2b256sum  string = "3c228306552177f5a304cb12a5b5e60897f2f486b64671afdccf0f8dd9410cbd"
)

func TestBLAKE2b256Message(t *testing.T) {
  test_blake2b256 := test.FoundBLAKE2b256(random_message)
  if blake2b256sum == test_blake2b256 {
    t.Log("blake2b256 found.")
  } else {
    t.Errorf("%v :: not found.", blake2b256sum)
  }
}

func TestBLAKE2s256Message(t *testing.T) {
  test_blake2s256 := test.FoundBLAKE2s256(random_message)
  if blake2s256sum == test_blake2s256 {
    t.Log("blake2s256 found.")
  } else {
    t.Errorf("%v :: not found.", blake2s256sum)
  }
}

func TestRIPEMD160Message(t *testing.T) {
  test_ripemd160 := test.FoundRIPEMD160(random_message)
  if ripemd160sum == test_ripemd160 {
    t.Log("ripemd160 found.")
  } else {
    t.Errorf("%v :: not found.", ripemd160sum)
  }
}

func TestSHA3_256Message(t *testing.T) {
  test_sha3_256 := test.FoundSHA3_256(random_message)
  if sha3_256sum == test_sha3_256 {
    t.Log("sha3_256 found.")
  } else {
    t.Errorf("%v :: not found.", sha3_256sum)
  }
}

func TestMD4Message(t *testing.T) {
  test_md4 := test.FoundMD4(random_message)
  if md4sum == test_md4 {
    t.Log("md4 found.")
  } else {
    t.Errorf("%v :: not found.", md4sum)
  }
}

func TestTIGER192Message(t *testing.T) {
  test_tiger192 := test.FoundTIGER192(random_message)
  if tiger192sum == test_tiger192 {
    t.Log("tiger192 found.")
  } else {
    t.Errorf("%v :: not found.", tiger192sum)
  }
}

func TestMD2Messagee(t *testing.T) {
  test_md2 := test.FoundMD2(random_message)
  if md2sum == test_md2 {
    t.Log("md2 found.")
  } else {
    t.Errorf("%v :: not found.", md2sum)
  }
}

func TestKECCAK256Message(t *testing.T) {
  test_keccak256 := test.FoundKECCAK256(random_message)
  if keccak256sum == test_keccak256 {
    t.Log("keccak256 found.")
  } else {
    t.Errorf("%v :: not found.", keccak256sum)
  }
}

func TestSHA256Message(t *testing.T) {
  test_sha256 := test.FoundSHA256(random_message)
  if sha256sum == test_sha256 {
    t.Log("sha256 found.")
  } else {
    t.Errorf("%v :: not found.", sha256sum)
  }
}
