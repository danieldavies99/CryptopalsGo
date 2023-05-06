package ascii

import (
	"bytes"
	"testing"
)

func TestBytesToAsciiString(t *testing.T) {
	got := BytesToAsciiString([]byte("Plant trees"))
	want := "Plant trees"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAsciiStringToBytes(t *testing.T) {
	got := AsciiStringToBytes("Plant trees")
	want := []byte("Plant trees")

	if !bytes.Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
