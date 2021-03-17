package main

import (
	"flag"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func usage() {
	usageError(nil)
}
func usageError(err error) {
	fmt.Fprintf(os.Stderr, "usage: crc32 [-decimal] [-polynomial ieee|castagnoli|koopman] [inputfile ...]\n")

	flag.PrintDefaults()
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nerror: %s\n", err)
	}
	os.Exit(2)
}

func printCrc32(filename string, content []byte, decimal bool, polynomial uint32, printFileName bool) {
	t := crc32.MakeTable(polynomial)
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

func polynomialNameToValue(name string) (uint32, error) {
	if strings.ToLower(name) == "ieee" {
		return crc32.IEEE, nil
	}
	if strings.ToLower(name) == "castagnoli" {
		return crc32.Castagnoli,  nil
	}
	if strings.ToLower(name) == "koopman" {
		return crc32.Koopman, nil
	}
	return crc32.IEEE, fmt.Errorf("%s is not a supported polynomial name", name)
}

func main() {
	flag.Usage = usage
	decimal := flag.Bool("decimal", false, "print crc32 in decimal notation")
	polynomialName := flag.String("polynomial", "castagnoli", "polynomial to use to calculate crc32: IEEE, Castagnoli or Koopman")
	flag.Parse()
	polynomial, err := polynomialNameToValue(*polynomialName)
	if err != nil {
		usageError(err)
	}

	args := flag.Args()
	if len(args) > 0 {
		for _, filename := range args {
			content, err := os.ReadFile(filename)
			if err != nil {
				log.Fatalf("[ERROR] failed to open file %s, %s", filename, err)
			}
			printCrc32(filename, content, *decimal, polynomial, true)
		}
	} else {
		content, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("[ERROR] failed to read from stdin, %s", err)
		}
		printCrc32("-", content, *decimal, polynomial, false)
	}
}
