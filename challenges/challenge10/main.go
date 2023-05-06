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
	fileContent, err := ioutil.ReadFile("challenge10.txt")
	if err != nil {
		log.Fatal(err)
	}

	unencodedFileContent, err := base64.Decode(fileContent)
	if err != nil {
		log.Fatal("Error decoding base64 string", err)
	}

	IV := []byte{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	decrypted, err := aes.DecryptAes128Cbc(
		ascii.AsciiStringToBytes("YELLOW SUBMARINE"),
		IV,
		unencodedFileContent,
	)
	if err != nil {
		log.Fatal("Error decrypting CBC", err)
	}
	fmt.Print("Decrypted: ", ascii.BytesToAsciiString(decrypted), "\n")
}
