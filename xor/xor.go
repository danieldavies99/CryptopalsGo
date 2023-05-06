package xor

import (
	"errors"
)

func XOR(input []byte, against []byte) ([]byte, error) {
	if len(input) != len(against) {
		return nil, errors.New("Cannot XOR byte slices with different lengths")
	}
	output := []byte{}

	for i := 0; i < len(input); i++ {
		output = append(output, input[i]^against[i])
	}

	return output, nil
}

func XORAgainstOneCharacter(input []byte, againstChar rune) ([]byte, error) {
	// Get byte string of againstChar
	against := []byte{}
	for i := 0; i < len(input); i++ {
		against = append(against, byte(againstChar))
	}

	return XOR(input, against)
}

func RepeatingKeyEncode(input []byte, key []byte) ([]byte, error) {
	output := []byte{}

	for index, character := range input {
		output = append(output, character^key[index%len(key)])
	}

	return output, nil
}


func RepeatingKeyDecode(input []byte, key []byte) ([]byte, error) {
	output := []byte{}

	for index, character := range input {
		output = append(output, character^key[index%len(key)])
	}

	return output, nil
}
