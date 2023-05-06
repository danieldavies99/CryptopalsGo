package xor

import (
	"bytes"
	"testing"

	"github.com/danieldavies99/cryptopals/hex"
)

func TestXOR(t *testing.T) {
	got, err := XOR(
		[]byte{28, 1, 17, 0, 31, 1, 1, 0, 6, 26, 2, 75, 83, 83, 80, 9, 24, 28},
		[]byte{104, 105, 116, 32, 116, 104, 101, 32, 98, 117, 108, 108, 39, 115, 32, 101, 121, 101},
	)
	if err != nil {
		t.Errorf("Got error %q", err)
	}
	want := []byte("the kid don't play")

	if !bytes.Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRepeatingKeyEncode(t *testing.T) {
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	got, err := RepeatingKeyEncode([]byte(input), []byte("ICE"))
	if err != nil {
		t.Errorf("Got error %q", err)
	}
	want, err := hex.HexStringToBytes("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")
	if err != nil {
		t.Errorf("Got error %q", err)
	}

	if !bytes.Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestRepeatingKeyD(t *testing.T) {
	input, err := hex.HexStringToBytes("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")
	got, err := RepeatingKeyEncode([]byte(input), []byte("ICE"))
	if err != nil {
		t.Errorf("Got error %q", err)
	}
	want := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")

	if !bytes.Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
