package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	hash, err := getSHA256Hash("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hash)
}

func getSHA256Hash(filepath string) (string, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:]), nil

}
