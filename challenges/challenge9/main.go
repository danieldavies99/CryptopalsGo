package main

import (
	"fmt"
	"log"

	"github.com/danieldavies99/cryptopals/ascii"
	"github.com/danieldavies99/cryptopals/padding"
)

func main() {
	input := ascii.AsciiStringToBytes("YELLOW SUBMARINE")
	padLength := 20

	paddedResult, err := padding.PKCS7Pad(input, padLength)
	if err != nil {
		log.Fatal("Error padding: ", err)
	}
	fmt.Print(paddedResult)
}
