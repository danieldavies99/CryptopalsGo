package hammingdistance

import (
	"bytes"
	"errors"

	"github.com/dgryski/go-bitstream"
)

func Calculate(input1 []byte, input2 []byte) (int, error) {
	if len(input1) != len(input2) {
		return 0, errors.New("Error calculating hamming distance, inputs do not have the same length")
	}

	input1Buffer := bytes.NewBuffer(input1)
	input2Buffer := bytes.NewBuffer(input2)

	input1bitstreamReader := bitstream.NewReader(input1Buffer)
	input2bitstreamReader := bitstream.NewReader(input2Buffer)

	result := 0
	for i := 0; i < len(input1)*8; i++ {
		bit1, err := input1bitstreamReader.ReadBit()
		bit2, err := input2bitstreamReader.ReadBit()
		if err != nil {
			return 0, err
		}
		if bit1 != bit2 {
			result += 1
		}
	}

	return result, nil
}
