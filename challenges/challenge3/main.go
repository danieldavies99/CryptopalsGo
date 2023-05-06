package main

import (
	"fmt"

	"github.com/danieldavies99/cryptopals/ascii"
	"github.com/danieldavies99/cryptopals/hex"
	"github.com/danieldavies99/cryptopals/scoretext"
	"github.com/danieldavies99/cryptopals/xor"
)

func main() {
	hexInput, err := hex.HexStringToBytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		fmt.Print("Error decoding hex ", err)
		return
	}

	scores := make(map[rune]float64)

	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	for _, char := range chars {
		xorRes, err := xor.XORAgainstOneCharacter(hexInput, char)
		if err != nil {
			fmt.Print("Error XOR ", err)
			return
		}
		score := scoretext.ScoreText(xorRes)
		scores[char] = score
	}

	highestScore := 0.0
	var highestScoreRune rune

	for char, score := range scores {
		if score > highestScore {
			highestScore = score
			highestScoreRune = char
		}
	}

	bestTextBytes, err := xor.XORAgainstOneCharacter(hexInput, highestScoreRune)
	bestTextString := ascii.BytesToAsciiString(bestTextBytes)
	fmt.Print("highest scoring rune ", string(highestScoreRune), "\nTranslated text: ", bestTextString)
}
