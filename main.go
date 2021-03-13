package main

import (
	"flag"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"log"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: crc32 [inputfile ...]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func printCrc32(filename string, content []byte, printFileName bool) {
	t := crc32.MakeTable(crc32.Castagnoli)
	r := crc32.Checksum(content, t)
	if printFileName {
		fmt.Printf("%d %s\n", r, filename)
	} else {
		fmt.Printf("%d\n", r)
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		for _, filename := range args {
			content, err := os.ReadFile(filename)
			if err != nil {
				log.Fatalf("[ERROR] failed to open file %s, %s", filename, err)
			}
			printCrc32(filename, content, true)
		}
	} else {
		content, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("[ERROR] failed to read from stdin, %s", err)
		}
		printCrc32("-", content, false)
	}
}
