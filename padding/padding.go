package padding

import (
	"errors"
)

func PKCS7Pad(input []byte, padLength int) ([]byte, error) {
	if padLength < len(input) {
		return nil, errors.New("Can't data that is shorter than the pad length")
	}

	lengthToPad := padLength - len(input)

	output := input
	for i := 0; i < lengthToPad; i++ {
		output = append(output, byte(lengthToPad))
	}
	return output, nil
}

func PKCS7PadToMultiple(input []byte, multiple int) ([]byte, error) {
	numToPad := multiple - (len(input) % multiple)
	if numToPad == multiple {
		numToPad = 0
	}
	paddedOutput, err := PKCS7Pad(input, len(input)+numToPad)
	if err != nil {
		return nil, err
	}
	return paddedOutput, nil
}
