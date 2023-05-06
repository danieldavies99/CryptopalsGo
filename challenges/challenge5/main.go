package main

import (
	"fmt"
	"log"

	"github.com/danieldavies99/cryptopals/ascii"
	"github.com/danieldavies99/cryptopals/hex"
	"github.com/danieldavies99/cryptopals/xor"
)

func main() {
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	repeatingKeyEncrypted, err := xor.RepeatingKeyEncode(
		ascii.AsciiStringToBytes(input),
		ascii.AsciiStringToBytes("ICE"),
	)
	if err != nil {
		log.Fatal(err)
	}

	hexOutput, err := hex.BytesToHexString(repeatingKeyEncrypted)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Ascii input:\n", input, "\n\nEncrypted output:\n", hexOutput)
}
