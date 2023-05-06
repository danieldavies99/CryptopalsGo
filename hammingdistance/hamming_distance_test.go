package hammingdistance

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	got, err := Calculate(
		[]byte("this is a test"),
		[]byte("wokka wokka!!!"),
	)
	if err != nil {
		t.Errorf("Got error %q", err)
	}
	want := 37

	if got != want {
		t.Errorf("got %q, wanted %q", fmt.Sprint(got), fmt.Sprint(want))
	}
}
