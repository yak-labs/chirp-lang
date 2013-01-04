package terp_test

import (
	. "terp"
	"testing"
)

func TestMatchTailStar(t *testing.T) {
	if !MatchTailStar("", "") {
		t.Errorf("chirp")
	}
	if !MatchTailStar("*", "") {
		t.Errorf("chirp")
	}
	if !MatchTailStar("*", "*") {
		t.Errorf("chirp")
	}
	if !MatchTailStar("*", "asdf") {
		t.Errorf("chirp")
	}
	if MatchTailStar("argle", "arglefargle") {
		t.Errorf("chirp")
	}
	if !MatchTailStar("argle*", "arglefargle") {
		t.Errorf("chirp")
	}
	if MatchTailStar("gle*", "arglefargle") {
		t.Errorf("chirp")
	}
	if MatchTailStar("fargle*", "arglefargle") {
		t.Errorf("chirp")
	}
}
