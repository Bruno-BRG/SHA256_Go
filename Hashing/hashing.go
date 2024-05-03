package main

/*teste*/

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

func getSHA256Hash(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:]), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		showMainMenu()
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		if choice == "3" {
			fmt.Println("Exiting...")
			os.Exit(0)
		}
		handleChoice(choice, reader)
	}
}

func showMainMenu() {
	fmt.Println("1. Hash a file")
	fmt.Println("2. Compare two hashes")
	fmt.Println("3. Exit")
	fmt.Print("Enter your choice: ")
}

func handleChoice(choice string, reader *bufio.Reader) {
	switch choice {
	case "1":
		fmt.Print("Enter the path to the file to hash: ")
		filePath, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		filePath = strings.TrimSpace(filePath)

		hash, err := getSHA256Hash(filePath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(hash)
	case "2":
		fmt.Print("Enter the first hash: ")
		hash1, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		hash1 = strings.TrimSpace(hash1)

		fmt.Print("Enter the second hash: ")
		hash2, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		hash2 = strings.TrimSpace(hash2)
		// Rest of the code...
	}
}
