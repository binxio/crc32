crc32 - calculate crc32 checksum
================================

## Usage
```
usage: crc32 [-decimal] [-polynomial ieee|castagnoli|koopman] [inputfile ...]

  -decimal
        print crc32 in decimal notation

  -polynomial string
        polynomial to use to calculate crc32: IEEE, Castagnoli or Koopman (default "castagnoli")

```

## Description

Calculate the crc32 checksum for the specified files. If not specified, it will
read the content from stdin. You can specify which polynomial to use. Cloud KMS uses
Castagnoli.

## Examples
to return crc32 checksum of specified files, type:
```sh
$ crc32 main.go README.md
aad0247b main.go
5cbc3315 README.md
```
to return crc32 checksum of stdin, type:

```sh
$ cat main.go | crc32
aad0247b
```
