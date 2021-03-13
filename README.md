crc32 - calculate crc32 checksum
================================

## Usage
```
	crc32 [-decimal] [ filename ...]
```

## Description

Calculate the crc32 checkfum for the specified files. If not specified, it will
read the content from stdin.

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
