crc32 - calculate crc32 checksum
================================

## Usage
```
	crc32 [ filename ...]
```

## Description

Calculate the crc32 checkfum for the specified files. If not specified, it will
read the content from stdin.

## Examples
to return crc32 checksum of specified files, type:
```sh
$ crc32 main.go README.md
1175033037 main.go
2279994761 README.md
```
to return crc32 checksum of stdin, type:

```sh
$ cat main.go | crc32
1175033037
```
