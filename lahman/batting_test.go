package lahman

import (
	"encoding/csv"
	"strings"
	"testing"
)

var linetwo string = "addybo01,1871,1,RC1,NA,25,118,30,32,6,0,0,13,8,1,4,0,,,,,"

// Test that reading from a single string parses correctly.
func Test_NewBatter(t *testing.T) {
	line, _ := csv.NewReader(strings.NewReader(linetwo)).Read()

	b, err := batter{}.csvRead(line)

	checkLineTwo(b.(*batter), err, t)

	if b.(*batter).String() != linetwo {
		t.Log(b)
		t.Log(linetwo)
		t.Fatal("batter line does not print properly")
	}
}

// Test that bad input returns an error message.
func Test_Parsing(t *testing.T) {
	linetwoerr := "addybo01,notayear,bob,RC1,NA,25,118,30,32,6,0,0,13,8,1,4,0,,,,,"

	line, _ := csv.NewReader(strings.NewReader(linetwoerr)).Read()
	_, err := batter{}.csvRead(line)

	if err == nil {
		t.Log(err)
		t.Fatal("Parsing failed to report error in 'stint'")
	}
}

// helper to test specific player line
func checkLineTwo(b *batter, err error, t *testing.T) {
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
	case b.H() != 32:
		t.Log("Hits: ", b.H())
		t.Fatal("Incorrectly parsed Hits")
		fallthrough
	case b.GIDP() != -1:
		t.Log("GIDP: ", b.GIDP())
		t.Fatal("Incorrectly parsed GIDP")
	}
}
