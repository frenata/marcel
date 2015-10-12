package lahman_test

// TODO: add tests for init() databases

import (
	"testing"

	"github.com/frenata/marcel/lahman"
)

var wiki2003 = "\nWiki Gonzalez,gonzawi01,2003,1,SDN,NL\n" +
	"Batting:  24,65,1,13,5,0,0,10,0,0,5,13,1,1,1,1,3\n" +
	"Pitching: 0,0,1,0,0,0,0,3,0,0,0,1,0,0.000,0.00,0,0,0,0,4,1,0,0,0,"

var tail string = "zuninmi01,2014,1,SEA,AL,131,438,51,87,20,2,22,60,0,3,17,158,1,17,0,4,12"
var pitchtail string = "zimmejo02,2014,1,WAS,NL,14,5,32,32,3,2,0,599,185,59,13,29,182,0.244,2.66,0,4,6,0,800,0,67,5,3,11"

func Test_GetYear(t *testing.T) {
	players := lahman.GetYear(2003)

	for _, p := range players {
		if p.IsPosPitcher() { // in 2003, only wiki
			if wiki2003 != players[439].String() {
				t.Log(players[439])
				t.Log(wiki2003)
				t.Fatal("Wiki Gonzalez should be the 439th entry in the database!")
			}
		}
	}
}

func Test_GetYear_PrintBatter(t *testing.T) {
	players := lahman.GetYear(2014)

	p := players[42] // Oswaldo Arcia

	switch {
	case p.Bat.PA() != 410:
		fallthrough
	case p.Bat.G() != 103:
		fallthrough
	case p.Bat.AB() != 372:
		fallthrough
	case p.Bat.R() != 46:
		fallthrough
	case p.Bat.H() != 86:
		fallthrough
	case p.Bat.H2() != 16:
		fallthrough
	case p.Bat.H3() != 3:
		fallthrough
	case p.Bat.HR() != 20:
		fallthrough
	case p.Bat.RBI() != 57:
		fallthrough
	case p.Bat.SB() != 1:
		fallthrough
	case p.Bat.CS() != 2:
		fallthrough
	case p.Bat.BB() != 31:
		fallthrough
	case p.Bat.SO() != 127:
		fallthrough
	case p.Bat.IBB() != 4:
		fallthrough
	case p.Bat.HBP() != 6:
		fallthrough
	case p.Bat.SH() != 0:
		fallthrough
	case p.Bat.SF() != 1:
		fallthrough
	case p.Bat.GIDP() != 6:
		t.Fatal("Error printing one of Oswaldo's stats")
	default:
		t.Log("Oswaldo printed fine.")
	}
}

func Test_GetYear_PrintPitcher(t *testing.T) {
	players := lahman.GetYear(2014)

	p := players[0] // Fernando Abad

	switch {
	case p.Pit.W() != 2:
		fallthrough
	case p.Pit.L() != 4:
		fallthrough
	case p.Pit.G() != 69:
		fallthrough
	case p.Pit.GS() != 0:
		fallthrough
	case p.Pit.CG() != 0:
		fallthrough
	case p.Pit.SHO() != 0:
		fallthrough
	case p.Pit.SV() != 0:
		fallthrough
	case p.Pit.IPouts() != 172:
		fallthrough
	case p.Pit.H() != 34:
		fallthrough
	case p.Pit.ER() != 10:
		fallthrough
	case p.Pit.HR() != 4:
		fallthrough
	case p.Pit.BB() != 15:
		fallthrough
	case p.Pit.SO() != 51:
		fallthrough
	case p.Pit.BAopp() != .175:
		fallthrough
	case p.Pit.ERA() != 1.57:
		fallthrough
	case p.Pit.IBB() != 3:
		fallthrough
	case p.Pit.WP() != 0:
		fallthrough
	case p.Pit.HBP() != 4:
		fallthrough
	case p.Pit.BK() != 0:
		fallthrough
	case p.Pit.BFP() != 216:
		fallthrough
	case p.Pit.GF() != 17:
		fallthrough
	case p.Pit.R() != 11:
		fallthrough
	case p.Pit.SH() != 1:
		fallthrough
	case p.Pit.SF() != 2:
		fallthrough
	case p.Pit.GIDP() != 6:
		t.Fatal("Error printing one of Abad's stats")
	default:
		t.Log("Abad printed fine.")
	}
}
