package main

import (
	"fmt"
	"log"

	"github.com/danieldavies99/cryptopals/ascii"
	"github.com/danieldavies99/cryptopals/hex"
	"github.com/danieldavies99/cryptopals/xor"
)

func main() {
	inputHexString := "1c0111001f010100061a024b53535009181c"
	inputAsBytes, err := hex.HexStringToBytes(inputHexString)
	if err != nil {
		log.Fatal(err)
	}

	againstHexString := "686974207468652062756c6c277320657965"
	againstAsBytes, err := hex.HexStringToBytes(againstHexString)
	if err != nil {
		log.Fatal(err)
	}

	xorRes, err := xor.XOR(inputAsBytes, againstAsBytes)
	if err != nil {
		log.Fatal(err)
	}

	xorAsHexString, err := hex.BytesToHexString(xorRes)
	if err != nil {
		log.Fatal(err)
	}
	// base64Encoded, err := base64.Encode(inputAsBytes)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Print(
		"Input hex: ", inputHexString,
		"\nAgainst hex: ", againstHexString,
		"\nRaw xor string: ", ascii.BytesToAsciiString(xorRes),
		"\nHex xor string: ", xorAsHexString)
}
