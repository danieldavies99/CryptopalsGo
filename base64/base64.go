package base64

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/dgryski/go-bitstream"
)

func removeElement(input []byte, val byte) []byte {
	j := 0
	for _, v := range input {
		if v != val {
			input[j] = v
			j++
		}
	}
	return input[:j]
}

func uintToB64Char(input uint64) (byte, error) {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	if input < 0 || int(input) > len(alphabet) {
		return 0, errors.New("Error converting int " + fmt.Sprint(input) + " to base64 character, out of range")
	}
	return alphabet[input], nil
}

func Encode(input []byte) ([]byte, error) {

	inputLengthInBits := len(input) * 8
	numToPad := 6 - inputLengthInBits%6
	if numToPad == 6 {
		numToPad = 0
	}
	inputLengthInBitsPlusPadding := inputLengthInBits + numToPad

	inputBuffer := bytes.NewBuffer(input)
	bitstreamWriter := bitstream.NewWriter(inputBuffer)
	// write padding bits
	for i := 0; i < numToPad; i++ {
		bitstreamWriter.WriteBit(bitstream.Bit(false)) // padding bit i.e. 0
	}
	bitstreamWriter.Flush(bitstream.Bit(false)) // can only write full bytes to stream so need to flush remaining

	// read bits into output
	output := []byte{}
	bitstreamReader := bitstream.NewReader(inputBuffer)

	for i := 0; i < inputLengthInBitsPlusPadding; i += 6 {
		nextByte, err := bitstreamReader.ReadBits(6)
		if err != nil {
			return nil, err
		}

		b64Char, err := uintToB64Char(nextByte)
		if err != nil {
			return nil, err
		}
		output = append(output, b64Char)
	}

	// add output padding
	numOuputPaddingChars := 3 - len(input)%3
	if numOuputPaddingChars == 3 {
		numOuputPaddingChars = 0
	}

	for i := 0; i < numOuputPaddingChars; i++ {
		output = append(output, 61)
	}

	return output, nil
}

func B64CharToUint(input byte) (uint64, error) {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	res := strings.Index(alphabet, string(input))
	return uint64(res), nil
}

func Decode(input []byte) ([]byte, error) {

	// nuke padding characters
	for input[len(input)-1] == '=' {
		input = input[:len(input)-1]
	}

	// nuke new lines
	input = removeElement(input, byte('\n'))

	outputBuffer := bytes.NewBuffer([]byte{})
	bitstreamWriter := bitstream.NewWriter(outputBuffer)

	totalBitsWritten := 0
	for i := 0; i < len(input); i++ {
		unencodedByte, err := B64CharToUint(input[i])
		if err != nil {
			return nil, err
		}
		bitstreamWriter.WriteBits(unencodedByte, 6)
		totalBitsWritten += 6
	}
	bitstreamWriter.Flush(bitstream.Bit(false))
	bitstreamReader := bitstream.NewReader(outputBuffer)

	output := []byte{}

	for i := 0; i < totalBitsWritten/8; i++ {
		currentByte, err := bitstreamReader.ReadByte()
		if err != nil {
			return nil, err
		}

		output = append(output, currentByte)
	}

	return output, nil
}
