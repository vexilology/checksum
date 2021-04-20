package algorithm

import (
  "hash"
  "math/rand"
  "crypto/sha1"
  "crypto/md5"
  "crypto/sha256"
  "encoding/base64"
  "testing"

  "golang.org/x/crypto/blake2b"
  "golang.org/x/crypto/blake2s"
  "golang.org/x/crypto/ripemd160"
  "golang.org/x/crypto/sha3"

  "github.com/vexilology/hashgosum/test"
)

type keccak256sum struct {
  hash   string
  result string
}

type sha256sum struct {
  hash   string
  result string
}

type md2sum struct {
  hash   string
  result string
}

type tiger192sum struct {
  hash   string
  result string
}

type md4sum struct {
  hash   string
  result string
}

type sha3_256sum struct {
  hash   string
  result string
}

type ripemd160sum struct {
  hash   string
  result string
}

type blake2s256sum struct {
  hash   string
  result string
}

type blake2b256sum struct {
  hash   string
  result string
}

var blake2b256_tests = []blake2b256sum{
  {"8928aae63c84d87ea098564d1e03ad813f107add474e56aedd286349c0c03ea4", "a"},
  {"bddd813c634239723171ef3fee98579b94964e3bb1cb3e427262c8c068d52319", "abc"},
  {"f5d67bae73b0e10d0dfd3043b3f4f100ada014c5c37bd5ce97813b13f5ab2bcf", "123"},
  {"3c228306552177f5a304cb12a5b5e60897f2f486b64671afdccf0f8dd9410cbd", "helloworld"},
  {"1fc70117dfe929b4c4f56e0629a6a6d36f86d8bcc5b7d705104e02e368485d8e", "hell@ #r1d"},
}

var blake2s256_tests = []blake2s256sum{
  {"4a0d129873403037c2cd9b9048203687f6233fb6738956e0349bd4320fec3e90", "a"},
  {"508c5e8c327c14e2e1a72ba34eeb452f37458b209ed63a294d999b4c86675982", "abc"},
  {"e906644ad861b58d47500e6c636ee3bf4cb4bb00016bb352b1d2d03d122c1605", "123"},
  {"909052ce75dc79b301b0347ac904bae3b3117ce5abf963f16ee4c8ce08a3d407", "helloworld"},
  {"6d899561b5d7d8a4a9a600b6ace446e9642e93567310f5950ad2f0e87eb0ba70", "hell@ #r1d"},
}

var ripemd160_tests = []ripemd160sum{
  {"0bdc9d2d256b3ee9daae347be6f4dc835a467ffe", "a"},
  {"8eb208f7e05d987a9b044a8e98c6b087f15a0bfc", "abc"},
  {"e3431a8e0adbf96fd140103dc6f63a3f8fa343ab", "123"},
  {"8a73c5438c28e79e696144fa869886f240cfaddb", "helloworld"},
  {"3b5f88c22833998ed456c3db9042b5ba3e0a23e1", "hell@ #r1d"},
}

var sha3_256tests = []sha3_256sum{
  {"80084bf2fba02475726feb2cab2d8215eab14bc6bdd8bfb2c8151257032ecd8b", "a"},
  {"3a985da74fe225b2045c172d6bd390bd855f086e3e9d525b46bfe24511431532", "abc"},
  {"a03ab19b866fc585b5cb1812a2f63ca861e7e7643ee5d43fd7106b623725fd67", "123"},
  {"92dad9443e4dd6d70a7f11872101ebff87e21798e4fbb26fa4bf590eb440e71b", "helloworld"},
  {"ab6145f1e068f73426aa12088a19e799a5349540c3c9c02e26a86fd8d75a04b3", "hell@ #r1d"},
}

var md4_tests = []md4sum{
  {"bde52cb31de33e46245e05fbdbd6fb24", "a"},
  {"a448017aaf21d8525fc10ae87aa6729d", "abc"},
  {"c58cda49f00748a3bc0fcfa511d516cb", "123"},
  {"793033db97268fc9ceebde269797e54b", "helloworld"},
  {"6c4a0c74f6bcd07045cdc8e71b8b9703", "hell@ #r1d"},
}

var tiger192_tests = []tiger192sum{
  {"77befbef2e7ef8ab2ec8f93bf587a7fc613e247f5f247809", "a"},
  {"2aab1484e8c158f2bfb8c5ff41b57a525129131c957b5f93", "abc"},
  {"a86807bb96a714fe9b22425893e698334cd71e36b0eef2be", "123"},
  {"5c9943da4a68b90dd5489ddef8d9eb6d71c9e0b00e2c4347", "helloworld"},
  {"d6ca216b4ad89f82c45c4bc955f7a4565cde22e35ac965ef", "hell@ #r1d"},
}

var md2_tests = []md2sum{
  {"32ec01ec4a6dac72c0ab96fb34c0b5d1", "a"},
  {"da853b0d3f88d99b30283a69e6ded6bb", "abc"},
  {"ef1fedf5d32ead6b7aaf687de4ed1b71", "123"},
  {"0ae06562456c6e0e9736f4feda3a477b", "helloworld"},
  {"9db878b5a03669b1b93270ae6d7fa7e8", "hell@ #r1d"},
}

var sha256_tests = []sha256sum{
  {"ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb", "a"},
  {"ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad", "abc"},
  {"a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3", "123"},
  {"936a185caaa266bb9cbe981e9e05cb78cd732b0b3280eb944412bb6f8f8f07af", "helloworld"},
  {"d316d2fd706f630fe0628642783851d2c49f76c1af6bd61dc5a827984382859a", "hell@ #r1d"},
}

var keccak256_tests = []keccak256sum{
  {"3ac225168df54212a25c1c01fd35bebfea408fdac2e31ddd6f80a4bbf9a5f1cb", "a"},
  {"4e03657aea45a94fc7d47ba826c8d667c0d1e6e33a64a036ec44f58fa12d6c45", "abc"},
  {"64e604787cbf194841e7b68d7cd28786f6c9a0a3ab9f8b0a0e87cb4387ab0107", "123"},
  {"fa26db7ca85ead399216e7c6316bc50ed24393c3122b582735e7f3b0f91b93f0", "helloworld"},
  {"25777fb75e8113f500c747c5a6f29203427293813d12963df6400897ce5fca3f", "hell@ #r1d"},
}

func benchmarkPrime(num uint, b *testing.B) {
  for i := 0; i < b.N; i++ {
    test.Prime(num)
  }
}

func benchmarkHash(h hash.Hash, b *testing.B) {
  data := make([]byte, 1024)
  rand.Read(data)

  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    h.Write(data)
    h.Sum(nil)
  }
}

func BenchmarkPrime1(b *testing.B) { benchmarkPrime(1, b) }

func BenchmarkPrime183(b *testing.B) { benchmarkPrime(183, b) }

func BenchmarkPrime923(b *testing.B) { benchmarkPrime(923, b) }

func BenchmarkPrime1039281(b *testing.B) { benchmarkPrime(1039281, b) }

func BenchmarkMD5(b *testing.B) { benchmarkHash(md5.New(), b) }

func BenchmarkSHA1(b *testing.B) { benchmarkHash(sha1.New(), b) }

func BenchmarkSHA256(b *testing.B) { benchmarkHash(sha256.New(), b) }

func BenchmarkSHA3_256(b *testing.B) { benchmarkHash(sha3.New256(), b) }

func BenchmarkRIPEMD160(b *testing.B) { benchmarkHash(ripemd160.New(), b) }

func BenchmarkBLAKE2b(b *testing.B) {
  h, _ := blake2b.New256(nil)
  benchmarkHash(h, b)
}

func BenchmarkBLAKE2s(b *testing.B) {
  h, _ := blake2s.New256(nil)
  benchmarkHash(h, b)
}

func BenchmarkEncode(b *testing.B) {
  data := make([]byte, 1024)
  rand.Read(data)

  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    base64.StdEncoding.EncodeToString([]byte(data))
  }
}

func BenchmarkDecode(b *testing.B) {
  data := make([]byte, 1024)
  rand.Read(data)
  encoded := base64.StdEncoding.EncodeToString([]byte(data))

  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    _, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
      panic(err)
    }
  }
}

func TestBLAKE2b256(t *testing.T) {
  for i := 0; i < len(blake2b256_tests); i++ {
    r := blake2b256_tests[i]
    f := test.FoundBLAKE2b256(r.result)
    if f != r.hash {
      t.Errorf("hash '%v' not found", r.hash)
    } else {
      t.Logf("message '%v' blake2b256sum '%v' found", r.result, r.hash)
    }
  }
}

func TestBLAKE2s256(t *testing.T) {
  for i := 0; i < len(blake2s256_tests); i++ {
    r := blake2s256_tests[i]
    f := test.FoundBLAKE2s256(r.result)
    if f != r.hash {
      t.Errorf("hash '%v' not found", r.hash)
    } else {
      t.Logf("message '%v' blake2s256sum '%v' found", r.result, r.hash)
    }
  }
}

func TestRIPEMD160(t *testing.T) {
  for i := 0; i < len(ripemd160_tests); i++ {
    r := ripemd160_tests[i]
    f := test.FoundRIPEMD160(r.result)
    if f != r.hash {
      t.Errorf("hash '%v' not found", r.hash)
    } else {
      t.Logf("message '%v' ripemd160sum '%v' found", r.result, r.hash)
    }
  }
}

func TestSHA3Message(t *testing.T) {
  for i := 0; i < len(sha3_256tests); i++ {
    r := sha3_256tests[i]
    f := test.FoundSHA3_256(r.result)
    if f != r.hash {
      t.Errorf("hash '%v' not found", r.hash)
    } else {
      t.Logf("message '%v' sha3_256sum '%v' found", r.result, r.hash)
    }
  }
}

func TestMD4Message(t *testing.T) {
  for i := 0; i < len(md4_tests); i++ {
    r := md4_tests[i]
    f := test.FoundMD4(r.result)
    if f != r.hash {
      t.Errorf("hash '%v' not found", r.hash)
    } else {
      t.Logf("message '%v' md4sum '%v' found", r.result, r.hash)
    }
  }
}

func TestTIGER192Message(t *testing.T) {
  for i := 0; i < len(tiger192_tests); i++ {
    r := tiger192_tests[i]
    f := test.FoundTIGER192(r.result)
    if f != r.hash {
      t.Errorf("hash '%v' not found", r.hash)
    } else {
      t.Logf("message '%v' tiger192sum '%v' found", r.result, r.hash)
    }
  }
}

func TestMD2Message(t *testing.T) {
  for i := 0; i < len(md2_tests); i++ {
    r := md2_tests[i]
    f := test.FoundMD2(r.result)
    if f != r.hash {
      t.Errorf("hash '%v' not found", r.hash)
    } else {
      t.Logf("message '%v' md2sum '%v' found", r.result, r.hash)
    }
  }
}

func TestKECCAK256Message(t *testing.T) {
  for i := 0; i < len(keccak256_tests); i++ {
    r := keccak256_tests[i]
    f := test.FoundKECCAK256(r.result)
    if f != r.hash {
      t.Errorf("hash '%v' not found", r.hash)
    } else {
      t.Logf("message '%v' keccak256sum '%v' found", r.result, r.hash)
    }
  }
}

func TestSHA256Message(t *testing.T) {
  for i := 0; i < len(sha256_tests); i++ {
    r := sha256_tests[i]
    f := test.FoundSHA256(r.result)
    if f != r.hash {
      t.Errorf("hash '%v' not found", r.hash)
    } else {
      t.Logf("message '%v' sha256sum '%v' found", r.result, r.hash)
    }
  }
}
