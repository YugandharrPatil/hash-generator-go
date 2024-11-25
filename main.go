package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"log"
	"os"
)

func main() {
	hashFunctionID := flag.Int(
		"function",
		0,
		`Which hashing function would you like to use? 
Select the hashing function using its integer ID: 
1: MD5
2: SHA-256
3: SHA-512
	`)

	dataToEncrypt := flag.String(
		"data",
		"",
		"the data you'd like to encrypt. don't wrap with quotation marks.",
	)

	flag.Parse()

	if len(os.Args) == 1 { // no flags passed
		flag.Usage()
		os.Exit(1)
	}

	if *dataToEncrypt == "" {
		log.Fatalln("Error: Please provide data using the -data flag.")
	}

	hashValue, err := encrypt(*hashFunctionID, *dataToEncrypt)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	fmt.Printf("Hash value: %x\n", hashValue)
}

func encrypt(id int, data string) ([]byte, error) {
	var h hash.Hash
	var functionName string

	switch id {
	case 1:
		h = md5.New()
		functionName = "MD5"
	case 2:
		h = sha256.New()
		functionName = "SHA-256"
	case 3:
		h = sha512.New()
		functionName = "SHA-512"
	default:
		return nil, fmt.Errorf(
			"invalid function id %d; valid options are 1 (MD5), 2 (SHA-256), 3 (SHA-512)", id,
		)
	}

	fmt.Printf("Using hash function: %s\n", functionName)
	if _, err := h.Write([]byte(data)); err != nil {
		return nil, fmt.Errorf("failed to write data to hash function %v", err)
	}
	return h.Sum(nil), nil
}
