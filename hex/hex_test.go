package hex

import (
	"bytes"
	"testing"
)

func TestHexStringToBytes(t *testing.T) {
	got, err := HexStringToBytes("506C616E74207472656573")
	if err != nil {
		t.Errorf("Got error %q", err)
	}
	want := []byte("Plant trees")

	if !bytes.Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestBytesToHexString(t *testing.T) {
	got, err := BytesToHexString([]byte("Plant trees"))
	if err != nil {
		t.Errorf("Got error %q", err)
	}
	want := "506C616E74207472656573"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
