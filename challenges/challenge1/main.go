package main

import (
	"fmt"
	"log"

	"github.com/danieldavies99/cryptopals/ascii"
	"github.com/danieldavies99/cryptopals/base64"
	"github.com/danieldavies99/cryptopals/hex"
)

func main() {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	inputAsBytes, err := hex.HexStringToBytes(hexString)
	if err != nil {
		log.Fatal(err)
	}

	base64Encoded, err := base64.Encode(inputAsBytes)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(
		"Input hex: ", hexString,
		"\nRaw string: ", ascii.BytesToAsciiString(inputAsBytes),
		"\nBase64 encoded string: ", ascii.BytesToAsciiString(base64Encoded))
}
