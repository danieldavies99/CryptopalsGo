package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/danieldavies99/cryptopals/aes"
	"github.com/danieldavies99/cryptopals/ascii"
	"github.com/danieldavies99/cryptopals/padding"
)

func generateRandomBytes(length int) []byte {
	res := make([]byte, length)
	rand.Read(res)
	return res
}

func appendRandomBytes(input []byte, numToAppend int) []byte {
	randomBytes := generateRandomBytes(numToAppend)
	return append(input, randomBytes...)
}

func prependRandomBytes(input []byte, numToPrepend int) []byte {
	randomBytes := generateRandomBytes(numToPrepend)
	return append(randomBytes, input...)
}

func encryptionOracle(input []byte) ([]byte, error) {
	// randomly pad input
	minPadding := 5
	maxPadding := 10
	// random number between minPadding and maxPadding
	numToPrepend := rand.Intn(maxPadding-minPadding) + minPadding
	numToAppend := rand.Intn(maxPadding-minPadding) + minPadding
	paddedInput := prependRandomBytes(input, numToPrepend)
	paddedInput = appendRandomBytes(paddedInput, numToAppend)

	// generate random key and random IV
	key := generateRandomBytes(16)
	IV := generateRandomBytes(16)

	shouldUseECB := rand.Intn(2) > 0

	if shouldUseECB {
		fmt.Print("Encrypting using ECB\n")
		blockPaddedInput, err := padding.PKCS7PadToMultiple(paddedInput, 16)
		if err != nil {
			return nil, err
		}
		encryptedOutput, err := aes.EncryptAes128Ecb(key, blockPaddedInput)
		if err != nil {
			return nil, err
		}
		return encryptedOutput, nil
	} else {
		fmt.Print("Encrypting using CBC\n")
		encryptedOutput, err := aes.EncryptAes128Cbc(key, IV, paddedInput)
		if err != nil {
			return nil, err
		}
		return encryptedOutput, nil
	}
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

func main() {
	rand.Seed(time.Now().UnixNano())
	// ECB detection only works if plain text has repetitions
	ptInput := ascii.AsciiStringToBytes(`AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA`)

	encrypted, err := encryptionOracle(ptInput)
	if err != nil {
		log.Fatal("Error encrypting using oracle, ", err)
	}

	repetitions := countRepetitions(encrypted, 16)
	fmt.Print("Repetitions: ", repetitions, "\n")

	if repetitions != 0 {
		fmt.Print("Detected encryption mode: ECB\n")
	} else {
		fmt.Print("Detected encryption mode: CBC\n")
	}
}
