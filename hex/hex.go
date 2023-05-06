package hex

import (
	"strconv"
	"strings"
)

func HexStringToBytes(input string) ([]byte, error) {
	output := []byte{}

	for i := 0; i < len(input); i += 2 {
		hexString := input[i : i+2]
		decimal, err := strconv.ParseInt(hexString, 16, 32)
		if err != nil {
			return nil, err
		}
		output = append(output, byte(decimal))
	}

	return output, nil
}

func BytesToHexString(input []byte) (string, error) {
	output := ""

	for i := 0; i < len(input); i++ {
		hexString := strings.ToUpper(strconv.FormatInt(int64(input[i]), 16))
		if len(hexString) < 2 {
			hexString = "0" + hexString
		}
		output += hexString
	}

	return output, nil
}
