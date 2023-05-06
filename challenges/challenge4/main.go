package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/danieldavies99/cryptopals/ascii"
	"github.com/danieldavies99/cryptopals/hex"
	"github.com/danieldavies99/cryptopals/scoretext"
	"github.com/danieldavies99/cryptopals/xor"
)

func main() {
	fileContent, err := ioutil.ReadFile("challenge4.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileContentString := string(fileContent)
	stringLines := strings.Split(fileContentString, "\n")
	lines := [][]byte{}

	for _, line := range stringLines {
		bytes, err := hex.HexStringToBytes(line)
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, bytes)
	}

	allChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	threshold := 2.0
	fmt.Print("Searching for text that scores above threshold: ", threshold, "\n\n")

	for lineIndex, line := range lines {
		for _, char := range allChars {
			xorRes, err := xor.XORAgainstOneCharacter(line, char)
			if err != nil {
				fmt.Print("Error XOR ", err)
				return
			}
			score := scoretext.ScoreText(xorRes)
			if score > threshold {
				highScoreRune := string(char)
				highScoreText := ascii.BytesToAsciiString(xorRes)
				fmt.Print("High scoring rune: ", highScoreRune, "\nScore: ", score, "\nFound on line: ", lineIndex, "\nTranslated text: ", highScoreText, "\n\n")
			}
		}
	}
}
