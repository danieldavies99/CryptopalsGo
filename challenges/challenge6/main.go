package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"github.com/danieldavies99/cryptopals/ascii"
	"github.com/danieldavies99/cryptopals/base64"

	"github.com/danieldavies99/cryptopals/hammingdistance"
	"github.com/danieldavies99/cryptopals/scoretext"
	"github.com/danieldavies99/cryptopals/xor"
)

func chunkByteSlice(input []byte, chunkSize int) [][]byte {
	output := [][]byte{}
	for i := 0; i < len(input)/(chunkSize+1); i++ {
		output = append(output, input[i*chunkSize:i*chunkSize+chunkSize])
	}
	return output
}

func transposeBlocks(input [][]byte) [][]byte {
	output := [][]byte{}
	for i := 0; i < len(input[0]); i++ {
		newBlock := []byte{}
		for j := 0; j < len(input); j++ {

			newBlock = append(newBlock, input[j][i])
		}
		output = append(output, newBlock)
	}
	return output
}

func getMostLikelySingleKeyChar(input []byte) rune {
	allChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789 :"

	bestScore := -1000000000.0 // arbitrarily low number
	bestRune := 'z'

	for _, char := range allChars {
		xorRes, err := xor.XORAgainstOneCharacter(input, char)
		if err != nil {
			log.Fatal("Error XOR ", err)
		}
		score := scoretext.ScoreText(xorRes)
		if score > bestScore {
			bestScore = score
			bestRune = char
		}
	}
	return bestRune
}

type distanceForKeysize struct {
	keysize  int
	distance float64
}

func getKeysizesByLiklihood(
	input []byte,
	minKeysize int,
	maxKeysize int,
) []distanceForKeysize {
	distancesByKeysize := []distanceForKeysize{}
	// find keysizes per distance
	for i := minKeysize; i <= maxKeysize; i++ {
		keysize := i
		totalBlocks := (len(input) / keysize) // total possible blocks

		distances := []int{}
		for j := 0; j < totalBlocks; j++ {
			firstKeysizeBytes := input[(j * keysize) : (j+1)*keysize]
			secondKeysizeBytes := input[(j+1)*keysize : (j+2)*keysize]
			distance, err := hammingdistance.Calculate(firstKeysizeBytes, secondKeysizeBytes)
			if err != nil {
				log.Fatal("error calculating hamming distance", err)
			}
			distances = append(distances, distance)
		}
		avgDistance := calculateAverage(distances)
		normalizedDistance := avgDistance / float64(keysize)
		fmt.Print(keysize, " ", normalizedDistance, "\n")
		distancesByKeysize = append(distancesByKeysize, distanceForKeysize{keysize, normalizedDistance})
	}

	// sort by distance ascending
	sort.Slice(distancesByKeysize, func(i, j int) bool {
		return distancesByKeysize[i].distance < distancesByKeysize[j].distance
	})

	return distancesByKeysize
}

func calculateAverage(input []int) float64 {
	// declaring a variable to store the sum
	sum := 0
	n := len(input)

	// traversing through the array using for loop
	for i := 0; i < n; i++ {
		// adding the values of array to the variable sum
		sum += (input[i])
	}
	return (float64(sum)) / (float64(n))
	// declaring a variable avg to find the average
}

func main() {
	fileContent, err := ioutil.ReadFile("challenge6.txt")
	if err != nil {
		log.Fatal(err)
	}

	unencodedFileContent, err := base64.Decode(fileContent)
	if err != nil {
		log.Fatal("Error decoding base64 string", err)
	}

	distancesByKeysize := getKeysizesByLiklihood(
		unencodedFileContent,
		2,
		40,
	)

	fmt.Print(distancesByKeysize, "\n")
	fmt.Print("best: ", distancesByKeysize[0], "\n")
	fmt.Print("second best: ", distancesByKeysize[1], "\n")
	fmt.Print("third best: ", distancesByKeysize[2], "\n\n")

	bestKeysizes := []int{
		distancesByKeysize[0].keysize,
		distancesByKeysize[1].keysize,
		distancesByKeysize[2].keysize,
	}

	for i := 0; i < len(bestKeysizes); i++ {
		keysize := bestKeysizes[i]
		blocks := chunkByteSlice(unencodedFileContent, keysize)

		transposedBlocks := transposeBlocks(blocks)

		key := []byte{}
		for j := 0; j < len(transposedBlocks); j++ {
			mostLikleyChar := getMostLikelySingleKeyChar(transposedBlocks[j])
			key = append(key, byte(mostLikleyChar))
		}

		fmt.Print("Key: ", string(key), "\n")
		res, err := xor.RepeatingKeyDecode(unencodedFileContent, key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("First 20 chars: ", ascii.BytesToAsciiString(res)[:20], "\n\n")
	}
}
