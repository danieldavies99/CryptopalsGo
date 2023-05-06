package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/danieldavies99/cryptopals/aes"
	"github.com/danieldavies99/cryptopals/ascii"
	"github.com/danieldavies99/cryptopals/base64"
)

func main() {
	fileContent, err := ioutil.ReadFile("challenge7.txt")
	if err != nil {
		log.Fatal(err)
	}

	unencodedFileContent, err := base64.Decode(fileContent)
	if err != nil {
		log.Fatal("Error decoding base64 string", err)
	}

	key := ascii.AsciiStringToBytes("YELLOW SUBMARINE")
	plainText, err := aes.DecryptAes128Ecb(key, unencodedFileContent)
	if err != nil {
		log.Fatal("Error decrypting AES 128ECB", err)
	}

	fmt.Print(ascii.BytesToAsciiString(plainText))
}
