package lahman_test

import (
	"encoding/csv"
	"strings"
	"testing"

	"github.com/frenata/marcel/lahman"
)

var linetwo string = "addybo01,1871,1,RC1,NA,25,118,30,32,6,0,0,13,8,1,4,0,,,,,"

// Test that creating a blank Batting profile works correctly.
func Test_BlankBatter(t *testing.T) {
	blank := struct {
		id          string
		year, stint int16
	}{"andrew", 2016, 0}

	b, _ := lahman.NewPlayer(blank.id, blank.year, blank.stint)

	switch {
	case b.ID != blank.id:
		t.Fatal("ID does not match")
	case b.Year != blank.year:
		t.Fatal("Year does not match")
	case b.Stint != blank.stint:
		t.Fatal("Stint does not match")
	}

}

// Test that reading from a single string parses correctly.
func Test_NewBatter(t *testing.T) {

	player, _ := csv.NewReader(strings.NewReader(linetwo)).Read()
	b, err := lahman.NewBatterCSV(player)

	checkLineTwo(b, err, t)
}

// Test that bad input returns an error message.
func Test_Parsing(t *testing.T) {
	linetwoerr := "addybo01,1871,bob,RC1,NA,25,118,30,32,6,0,0,13,8,1,4,0,,,,,"

	player, _ := csv.NewReader(strings.NewReader(linetwoerr)).Read()
	_, err := lahman.NewBatterCSV(player)

	if err == nil {
		t.Log(err)
		t.Fatal("Parsing failed to report error in 'stint'")
	}
}

// helper to test specific player line
func checkLineTwo(b *lahman.Batter, err error, t *testing.T) {
	switch {
	case err != nil:
		t.Log(err)
		t.Fatal("Parsing error")
	case b.ID != "addybo01":
		t.Log("ID: ", b.ID)
		t.Fatal("Incorrectly parsed ID")
		fallthrough
	case b.League != "NA":
		t.Log("League: ", b.League)
		t.Fatal("Incorrectly parsed League")
		fallthrough
	case b.H != 32:
		t.Log("Hits: ", b.H)
		t.Fatal("Incorrectly parsed Hits")
		fallthrough
	case b.GIDP != 0:
		t.Log("GIDP: ", b.GIDP)
		t.Fatal("Incorrectly parsed GIDP")
	}
}
