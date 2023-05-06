package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/danieldavies99/cryptopals/aes"
	"github.com/danieldavies99/cryptopals/ascii"
	"github.com/danieldavies99/cryptopals/base64"
	"github.com/danieldavies99/cryptopals/padding"
)

func encryptionOracle(input []byte) ([]byte, error) {
	// generate random key and random IV
	key := ascii.AsciiStringToBytes("YELLOW SUBMARINE")

	// Encrypt with ECB
	blockPaddedInput, err := padding.PKCS7PadToMultiple(input, 16)
	if err != nil {
		return nil, err
	}
	encryptedOutput, err := aes.EncryptAes128Ecb(key, blockPaddedInput)
	if err != nil {
		return nil, err
	}
	return encryptedOutput, nil
}

func getAllSubstringsOfSize(input []byte, chunkSize int) [][]byte {
	totalSubstrings := len(input) - chunkSize
	output := [][]byte{}
	for i := 0; i < totalSubstrings; i++ {
		output = append(output, input[i:i+chunkSize])
	}
	return output
}

func countRepetitions(
	input []byte,
	blockSize int,
) int {
	blocks := getAllSubstringsOfSize(input, blockSize)
	blockHistory := [][]byte{}
	duplicateCount := 0
	for _, block := range blocks {
		isBlockDuplicate := false
		for _, alreadyFoundBlock := range blockHistory {
			if bytes.Equal(block, alreadyFoundBlock) {
				isBlockDuplicate = true
				duplicateCount++
			}
		}
		if !isBlockDuplicate {
			blockHistory = append(blockHistory, block)
		}
	}
	return duplicateCount
}

func calculateBlocksize() int {
	testBytes := []byte{'A'}
	initialEncryption, _ := encryptionOracle(testBytes)

	initialLength := len(initialEncryption)
	currentLength := initialLength
	for initialLength == currentLength {
		testBytes = append(testBytes, 'A')
		encrypted, _ := encryptionOracle(testBytes)
		currentLength = len(encrypted)
	}
	return currentLength - initialLength
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// ECB detection only works if plain text has repetitions
	ptInput := ascii.AsciiStringToBytes(
		`AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA`,
	)

	unknownString, err := base64.Decode(
		ascii.AsciiStringToBytes(`Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK`),
	)
	if err != nil {
		log.Fatal("Error decoding base64, ", err)
	}

	combinedInput := append(ptInput, unknownString...)

	encrypted, err := encryptionOracle(combinedInput)
	if err != nil {
		log.Fatal("Error encrypting using oracle, ", err)
	}

	repetitions := countRepetitions(encrypted, 16)
	fmt.Print("Repetitions: ", repetitions, "\n")

	if repetitions != 0 {
		fmt.Print("Detected encryption mode: ECB\n")
	} else {
		log.Fatal("Detected encryption mode: CBC\n")
	}

	// calculate block size
	blockSize := calculateBlocksize()
	fmt.Print("Detected block size ", blockSize, "\n")
}
