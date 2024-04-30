package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("1. Hash a file")
	fmt.Println("2. Compare two hashes")
	fmt.Print("Enter your choice: ")

	choice, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Print("Enter the path to the file to hash: ")
		filePath, _ := reader.ReadString('\n')
		filePath = strings.TrimSpace(filePath)

		hash, err := getSHA256Hash(filePath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(hash)
	case "2":
		fmt.Print("Enter the first hash: ")
		hash1, _ := reader.ReadString('\n')
		hash1 = strings.TrimSpace(hash1)

		fmt.Print("Enter the second hash: ")
		hash2, _ := reader.ReadString('\n')
		hash2 = strings.TrimSpace(hash2)

		if hash1 == hash2 {
			fmt.Println("The hashes match.")
		} else {
			fmt.Println("The hashes do not match.")
		}
	default:
		fmt.Println("Invalid choice.")
	}
}

func getSHA256Hash(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:]), nil
}
