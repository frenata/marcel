package lahman_test

import (
	"fmt"
	"testing"

	"github.com/frenata/marcel/lahman"
)

var tail string = "zuninmi01,2014,1,SEA,AL,131,438,51,87,20,2,22,60,0,3,17,158,1,17,0,4,12"

// Test reading from the CSV database
func Test_ParseCSV(t *testing.T) {
	r, err := lahman.ReadCSV("data/Batting.csv")
	if err != nil {
		t.Log(err)
		t.Fatal("error reading file")
	}

	player, _ := r.Read() // dispose of first line
	player, _ = r.Read()  // dispose of second line
	player, _ = r.Read()

	b, err := lahman.NewBatterCSV(player)
	checkLineTwo(b, err, t)
}

// test csv read when given bad file name
func Test_BadFile(t *testing.T) {
	_, err := lahman.ReadCSV("data/bad.csv")
	if err == nil {
		t.Fatal("failed to report error on bad file")
	}
}

// test read full file, check last line
func Test_ReadFull(t *testing.T) {
	batters, err := lahman.ReadAll("data/Batting.csv")
	if err != nil {
		t.Fatal("error reading full file")
	}
	fmt.Println(batters[len(batters)-1])
}
