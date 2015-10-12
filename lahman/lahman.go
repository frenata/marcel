package lahman

import (
	"log"
	"os"
)

var bats string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/BattingPost.csv"
var pits string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/PitchingPost.csv"

var batDB []*batter
var pitDB []*pitcher

// GetYear returns a list of all Players, with batting and pitching lines, for that year.
func GetYear(year float64) []*Player {
	batters := battingYear(year)
	pitchers := pitchingYear(year)
	players := make([]*Player, len(batters))

	for i, b := range batters {
		players[i] = &Player{}
		players[i].b = true
		players[i].bio = b.bio
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
	lines, err := ReadAll(bats, batter{})
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Batters loaded", len(lines))
	res := []*batter{}
	for _, l := range lines {
		b := l.(*batter)
		res = append(res, b)
	}
	batDB = res

	plines, err := ReadAll(pits, pitcher{})
	//fmt.Println("Pitchers loaded", len(plines))
	if err != nil {
		log.Fatal(err)
	}
	pres := []*pitcher{}
	for _, l := range plines {
		b := l.(*pitcher)
		pres = append(pres, b)
	}
	pitDB = pres
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
