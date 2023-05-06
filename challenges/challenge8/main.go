package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"

	"github.com/danieldavies99/cryptopals/hex"
)

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
	fileContent, err := ioutil.ReadFile("challenge8.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileContentAsString := string(fileContent)
	lines := strings.Split(fileContentAsString, "\n")

	type repetitionsForLine struct {
		line        int
		repetitions int
	}

	repetitionsForLineResult := []repetitionsForLine{}

	for i, line := range lines {
		lineData, err := hex.HexStringToBytes(line)
		if err != nil {
			log.Fatal("Error decoding hex ", err)
		}
		repetitions := countRepetitions(lineData, 16)
		repetitionsForLineResult = append(repetitionsForLineResult, repetitionsForLine{i, repetitions})
	}

	sort.Slice(repetitionsForLineResult, func(i, j int) bool {
		return repetitionsForLineResult[i].repetitions > repetitionsForLineResult[j].repetitions
	})

	fmt.Print(
		"Line ",
		repetitionsForLineResult[0].line,
		" has the most 16 byte block repetitions with ",
		repetitionsForLineResult[0].repetitions,
		" repeitions making it the most likely to be AES ECB envrypted",
	)
}
