package scoretext

import (
	"testing"
)

func TestScoreText(t *testing.T) {
	goodScore := ScoreText([]byte("Hi, I'm sensible english"))
	badScore := ScoreText([]byte("ヽ༼ ຈل͜ຈ༼ ▀̿̿Ĺ̯̿̿▀̿ ̿༽Ɵ͆ل͜Ɵ͆ ༽ﾉ"))

	if goodScore < badScore {
		t.Errorf("Got good score of %q, which is less than bad score of %q", goodScore, badScore)
	}
}
