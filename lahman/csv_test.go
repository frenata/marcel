package lahman

import (
	"fmt"
	"testing"
)

var tail string = "zuninmi01,2014,1,SEA,AL,131,438,51,87,20,2,22,60,0,3,17,158,1,17,0,4,12"

func Test_PitchPrint(t *testing.T) {
	r, err := newReader("data/Pitching.csv")
	if err != nil {
		t.Log(err)
		t.Fatal("error reading file")
	}
	line, _ := r.Read()
	line, _ = r.Read()
	line, _ = r.Read()

	p, err := Pitcher{}.csvRead(line)
	_ = fmt.Sprint(p)
}

// Test reading from the CSV database
func Test_ParseCSV(t *testing.T) {
	r, err := newReader("data/Batting.csv")
	if err != nil {
		t.Log(err)
		t.Fatal("error reading file")
	}

	player, _ := r.Read() // dispose of first line
	player, _ = r.Read()  // dispose of second line
	player, _ = r.Read()

	b, err := Batter{}.csvRead(player)
	checkLineTwo(b.(*Batter), err, t)
}

// test csv read when given bad file name
func Test_BadFile(t *testing.T) {
	_, err := newReader("data/bad.csv")
	if err == nil {
		t.Fatal("failed to report error on bad file")
	}
}

// test read full file, check last line
func Test_ReadFull(t *testing.T) {
	batters, err := ReadAll("data/Batting.csv")
	if err != nil {
		t.Fatal("error reading full file")
	}
	//	fmt.Println(batters[len(batters)-1])

	b := batters[len(batters)-1].(*Batter)
	// add checks for last batter
	switch {
	case err != nil:
		t.Log(err)
		t.Fatal("Parsing error")
	case b.ID != "zuninmi01":
		t.Log("ID: ", b.ID)
		t.Fatal("Incorrectly parsed ID")
		fallthrough
	case b.Team != "SEA":
		t.Log("Team: ", b.Team)
		t.Fatal("Incorrectly parsed Team")
		fallthrough
	case b.HR != 22:
		t.Log("Homeruns: ", b.HR)
		t.Fatal("Incorrectly parsed HR")
		fallthrough
	case b.SF != 4:
		t.Log("SF: ", b.SF)
		t.Fatal("Incorrectly parsed SF")
	}
}
