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
	fmt.Fprintf(os.Stderr, "usage: crc32 [-decimal] [inputfile ...]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func printCrc32(filename string, content []byte, decimal bool, printFileName bool) {
	t := crc32.MakeTable(crc32.Castagnoli)
	crc := crc32.Checksum(content, t)

	if decimal {
		fmt.Printf("%d", crc)
	} else {
		fmt.Printf("%x", crc)
	}
	if printFileName {
		fmt.Printf(" %s", filename)
	}
	fmt.Printf("\n")
}

func main() {
	flag.Usage = usage
	decimal := flag.Bool("decimal", false, "print crc32 in decimal notation")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		for _, filename := range args {
			content, err := os.ReadFile(filename)
			if err != nil {
				log.Fatalf("[ERROR] failed to open file %s, %s", filename, err)
			}
			printCrc32(filename, content, *decimal, true)
		}
	} else {
		content, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("[ERROR] failed to read from stdin, %s", err)
		}
		printCrc32("-", content, *decimal, false)
	}
}
