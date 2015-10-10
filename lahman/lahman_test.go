package lahman_test

import (
	"testing"

	"github.com/frenata/marcel/lahman"
)

var tail string = "zuninmi01,2014,1,SEA,AL,131,438,51,87,20,2,22,60,0,3,17,158,1,17,0,4,12"

func Test_BattingYear(t *testing.T) {
	batters, _ := lahman.ReadAll("data/Batting.csv", lahman.Batter{})
	res, _ := lahman.BattingYear(2014, batters)

	if len(res) != 1435 {
		t.Fatal("length of 2014 batters is not 1435")
	}

	if res[len(res)-1].String() != tail {
		t.Fatal("Last record does not match tail")
	}
}
