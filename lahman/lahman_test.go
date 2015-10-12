package lahman_test

import (
	"testing"

	"github.com/frenata/marcel/lahman"
)

var tail string = "zuninmi01,2014,1,SEA,AL,131,438,51,87,20,2,22,60,0,3,17,158,1,17,0,4,12"
var pitchtail string = "zimmejo02,2014,1,WAS,NL,14,5,32,32,3,2,0,599,185,59,13,29,182,0.244,2.66,0,4,6,0,800,0,67,5,3,11"

func Test_GetYear(t *testing.T) {
	lahman.GetYear(2003)
}

/*
func Test_BattingYear(t *testing.T) {
	res := lahman.battingYear(2014)

	if len(res) != 1435 {
		t.Fatal("length of 2014 batters is not 1435")
	}

	if res[len(res)-1].String() != tail {
		t.Log(res[len(res)-1])
		t.Log(tail)
		t.Fatal("Last record does not match tail")
	}
}

func Test_PitchingYear(t *testing.T) {
	res := lahman.pitchingYear(2014)

	if len(res) != 746 {
		t.Log(len(res))
		t.Fatal("length of 2014 batters is not 1435")
	}

	if res[len(res)-1].String() != pitchtail {
		t.Log(res[len(res)-1])
		t.Log(pitchtail)
		t.Fatal("Last record does not match tail")
	}
}
*/
