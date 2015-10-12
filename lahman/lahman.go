package lahman

import (
	"log"
	"os"
)

var master string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/Master.csv"
var batR string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/Batting.csv"
var pitR string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/Pitching.csv"
var batP string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/BattingPost.csv"
var pitP string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/PitchingPost.csv"

var batPDB []*batter
var pitPDB []*pitcher
var batDB []*batter
var pitDB []*pitcher
var masDB map[string]Master

// GetYear returns a list of all Players, with batting and pitching lines, for that year.
func GetYear(year float64) []*Player {
	batters := battingYear(year)
	pitchers := pitchingYear(year)
	players := make([]*Player, len(batters))

	for i, b := range batters {
		players[i] = &Player{}
		players[i].b = true
		players[i].bio = b.bio
		players[i].Mas = masDB[b.bio.ID]
		players[i].bio.First = masDB[b.bio.ID][13]
		players[i].bio.Last = masDB[b.bio.ID][14]
		players[i].Bat = b.BatStats
		for _, p := range pitchers {
			if p.bio == b.bio {
				players[i].Pit = p.PitchStats
				players[i].p = true
				if players[i].Pit.BFP() > players[i].Bat.AB() {
					players[i].b = false
				}
				break
			}
		}
	}
	return players
}

// initializes the pitching and batting databases into package variables
func init() {
	// init the batter database
	blines, err := ReadAll(batR, batter{})
	if err != nil {
		log.Fatal(err)
	}
	bres := []*batter{}
	for _, l := range blines {
		b := l.(*batter)
		bres = append(bres, b)
	}
	batDB = bres

	// init the pitcher db
	plines, err := ReadAll(pitR, pitcher{})
	if err != nil {
		log.Fatal(err)
	}
	pres := []*pitcher{}
	for _, l := range plines {
		b := l.(*pitcher)
		pres = append(pres, b)
	}
	pitDB = pres

	// init the postseason batter database
	blines, err = ReadAll(batP, batter{})
	if err != nil {
		log.Fatal(err)
	}
	bres = []*batter{}
	for _, l := range blines {
		b := l.(*batter)
		bres = append(bres, b)
	}
	batPDB = bres

	// init the postseason pitcher db
	plines, err = ReadAll(pitP, pitcher{})
	if err != nil {
		log.Fatal(err)
	}
	pres = []*pitcher{}
	for _, l := range plines {
		b := l.(*pitcher)
		pres = append(pres, b)
	}
	pitPDB = pres

	// init the master database
	masDB = make(map[string]Master)
	mlines, err := ReadAll(master, Master{})
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range mlines {
		m := l.(Master)
		masDB[m[0]] = m
		//pres = append(pres, b)
	}
}

// read the batting database and returns a list of batters for a given year
func battingYear(year float64) []*batter {
	res := []*batter{}
	for _, b := range batDB {
		if b.Year == year {
			res = append(res, b)
		}
	}
	return res
}

// read the pitching database and return a list of pitchers for a given year
func pitchingYear(year float64) []*pitcher {
	res := []*pitcher{}
	for _, p := range pitDB {
		if p.Year == year {
			res = append(res, p)
		}
	}
	return res
}
