> Testing on Ubuntu 16.04 && go1.13.4 linux/amd64.

## Start

```
# need golang library crypto and other packages
make install
```

```
# now you can use it
make help
make build

hashgosum -m
hashgosum -f=FILENAME

or

go run main.go -m
go run main.go -f=FILENAME
```

## Binary

| Algorithm |
| --------- |
| ascii85	  |
| base32    |
| base64    |

## Hash algorithms

| Algorithm       | Digest sizes  |
| --------------- | ------------- |
| tiger192,3      |  192 - bit    |
| shake128-256    |  256 - bit    |
| shake256-512    |  512 - bit    |
| keccak224       |  224 - bit    |
| keccak256       |  256 - bit    |
| keccak384       |  384 - bit    |
| keccak512       |  512 - bit    |
| crc32-ieee      |  32  - bit    |
| crc64-ecma      |  64  - bit    |
| crc64-iso       |  64  - bit    |
| adler32         |  32  - bit    |
| blake2s-256     |  256 - bit    |
| blake2b-256     |  256 - bit    |
| blake2b-384     |  384 - bit    |
| blake2b-512     |  512 - bit    |
| ripemd160       |  160 - bit    |
| md2             |  128 - bit    |
| md4             |  128 - bit    |
| md5             |  128 - bit    |
| sha1            |  160 - bit    |
| sha224          |  224 - bit    |
| sha256          |  256 - bit    |
| sha384          |  384 - bit    |
| sha512          |  512 - bit    |
| sha512-224      |  224 - bit    |
| sha512-256      |  256 - bit    |
| sha3-224        |  224 - bit    |
| sha3-256        |  256 - bit    |
| sha3-384        |  384 - bit    |
| sha3-512        |  512 - bit    |
| fnv1-32         |  32  - bit    |
| fnv1-64         |  64  - bit    |
| fnv1-128        |  128 - bit    |
| fnv1a-32        |  32  - bit    |
| fnv1a-64        |  64  - bit    |
| fnv1a-128       |  128 - bit    |

