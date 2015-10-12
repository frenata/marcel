package lahman

import (
	"encoding/csv"
	"strings"
	"testing"
)

var pitchtwo string = "brainas01,1871,1,WS3,NA,12,15,30,30,30,0,0,792,361,132,4,37,13,,4.5,,,,0,,,292,,,"

// Test that reading from a single string parses correctly.
func Test_NewPitcher(t *testing.T) {

	line, _ := csv.NewReader(strings.NewReader(pitchtwo)).Read()

	p, err := pitcher{}.csvRead(line)
	if err != nil {
		t.Log(err)
		t.Fatal("error parsing Pticher")
	}

	checkPitchTwo(p.(*pitcher), err, t)
}

// Test that bad input returns an error message.
func Test_PitchParsing(t *testing.T) {
	pitchtwoerr := "brainas01,18.71,1,WS3,NA,12,biff,30,30,30,0,0,792,361,132,4,37,13,,4.5,,,,0,,,292,,,"

	line, _ := csv.NewReader(strings.NewReader(pitchtwoerr)).Read()
	_, err := pitcher{}.csvRead(line)

	if err == nil {
		t.Log(err)
		t.Fatal("Parsing failed to report error in 'stint'")
	}
}

// helper to test specific player line
func checkPitchTwo(p *pitcher, err error, t *testing.T) {
	switch {
	case err != nil:
		t.Log(err)
		t.Fatal("Parsing error")
	case p.ID != "brainas01":
		t.Log("ID: ", p.ID)
		t.Fatal("Incorrectly parsed ID")
		fallthrough
	case p.League != "NA":
		t.Log("League: ", p.League)
		t.Fatal("Incorrectly parsed League")
		fallthrough
	case p.H() != 361:
		t.Log("Hits: ", p.H())
		t.Fatal("Incorrectly parsed Hits")
		fallthrough
	case p.GIDP() != -1:
		t.Log("GIDP: ", p.GIDP())
		t.Fatal("Incorrectly parsed GIDP")
	}
}
