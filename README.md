# hashgosum

> Testing on Ubuntu 16.04 && go1.13.4 linux/amd64.

> Testing on Ubuntu 20.04.1 LTS on Windows 10 x86_64 && go1.15.2 linux/amd64.

## First Start

```
go get -u github.com/vexilology/hashgosum
cd ~/go/bin
chmod +x hashgosum
sudo ln -sf ~/go/bin/hashgosum /usr/bin/hashgosum

or

git clone git@github.com:vexilology/hashgosum.git
cd hashgosum/
make
```

## Available commands
```
make list
```

## Examples
```
hashgosum -h
hashgosum -s="hello World" -a=sha256
hashgosum -f=filename -a=sha256
```

## Hash algorithms
<details>
  <summary>Expand to view</summary>

| Algorithm       | Digest sizes  |
| --------------- | ------------- |
| tiger192,3      |  192 - bit    |
| shake128-224    |  224 - bit    |
| shake128-256    |  256 - bit    |
| shake128-384    |  384 - bit    |
| shake128-512    |  512 - bit    |
| shake256-224    |  224 - bit    |
| shake256-256    |  256 - bit    |
| shake256-384    |  384 - bit    |
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

</details>
