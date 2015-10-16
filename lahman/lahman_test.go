package lahman_test

// TODO: add tests for init() databases

import (
	"testing"

	"github.com/frenata/marcel/lahman"
)

var wiki2003 = "Wiki Gonzalez,gonzawi01,2003,1,SDN,NL\n" +
	"Batting:  24,65,1,13,5,0,0,10,0,0,5,13,1,1,1,1,3\n" +
	"Pitching: 0,0,1,0,0,0,0,3,0,0,0,1,0,0.000,0.00,0,0,0,0,4,1,0,0,0,"

var tail string = "zuninmi01,2014,1,SEA,AL,131,438,51,87,20,2,22,60,0,3,17,158,1,17,0,4,12"
var pitchtail string = "zimmejo02,2014,1,WAS,NL,14,5,32,32,3,2,0,599,185,59,13,29,182,0.244,2.66,0,4,6,0,800,0,67,5,3,11"

func Test_GetYear(t *testing.T) {
	players := lahman.GetYear(2003)

	none := true
	for _, p := range players {
		if p.IsPosPitcher() { // in 2003, only wiki
			none = false
			if wiki2003 != p.String() {
				t.Log(p)
				t.Log(wiki2003)
				t.Fatal("Wiki Gonzalez should be the 439th entry in the database!")
			}
		}
	}
	if none {
		t.Fatal("Wiki isn't there! :(")
	}
}

func Test_GetPostYear(t *testing.T) {
	players := lahman.GetPostYear(2010)

	none := true
	for _, p := range players {
		if p.Pit.SHO() > 0 && p.Pit.H() == 0 { // Roy Halladay 2010
			none = false
			if p.Year() != 2010 {
				t.Log(p.Year())
				t.Fatal("Year should be 2010")
			}
			if p.Stint() != "NLDS1" {
				t.Log(p.Stint())
				t.Fatal("Stint should be NLDS1")
			}
		}
	}
	if none {
		t.Fatal("Halladay isn't there")
	}
}
