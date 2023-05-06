package base64

import (
	"bytes"
	"testing"
)

func TestEncodeOnePaddingChar(t *testing.T) {
	got, err := Encode([]byte("light work."))
	if err != nil {
		t.Errorf("Got error %q", err)
	}
	want := []byte("bGlnaHQgd29yay4=")

	if !bytes.Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestEncodeTwoPaddingChars(t *testing.T) {
	got, err := Encode([]byte("light work"))
	if err != nil {
		t.Errorf("Got error %q", err)
	}
	want := []byte("bGlnaHQgd29yaw==")

	if !bytes.Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestEncodeNoPaddingChars(t *testing.T) {
	got, err := Encode([]byte("light wor"))
	if err != nil {
		t.Errorf("Got error %q", err)
	}
	want := []byte("bGlnaHQgd29y")

	if !bytes.Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDecode(t *testing.T) {
	got, err := Decode([]byte("bGlnaHQgd29yay4="))
	if err != nil {
		t.Errorf("Got error %q", err)
	}
	want := []byte("light work.")

	if !bytes.Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDecodeNoPadding(t *testing.T) {
	got, err := Decode([]byte("U29tZWJvZHkgb25jZSB0b2xkIG1lIHRoZSB3b3JsZCBpcyBnb25uYSByb2xsIG1lCkkgYWluJ3QgdGhlIHNoYXJwZXN0IHRvb2wgaW4gdGhlIHNoZWQKU2hlIHdhcyBsb29raW5nIGtpbmQgb2YgZHVtYiB3aXRoIGhlciBmaW5nZXIgYW5kIGhlciB0aHVtYgpJbiB0aGUgc2hhcGUgb2YgYW4gIkwiIG9uIGhlciBmb3JlaGVhZA=="))
	if err != nil {
		t.Errorf("Got error %q", err)
	}
	want := []byte("Somebody once told me the world is gonna roll me\nI ain't the sharpest tool in the shed\nShe was looking kind of dumb with her finger and her thumb\nIn the shape of an \"L\" on her forehead")

	if !bytes.Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
