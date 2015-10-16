package lahman

import (
	"log"
	"os"
)

type db struct {
	years        []int
	batPostDB    batDB
	pitPostDB    pitDB
	batRegularDB batDB
	pitRegularDB pitDB
	masterDB     map[string]master
}

type batDB []*batter
type pitDB []*pitcher

var cur db

var (
	masterCsv    string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/Master.csv"
	batting      string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/Batting.csv"
	pitching     string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/Pitching.csv"
	battingPost  string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/BattingPost.csv"
	pitchingPost string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/PitchingPost.csv"
)

// GetPlayer returns the past 3 years of data for that Player
// TODO: test
func GetPlayer(id string, year int, n int) []*Player {
	results := []*Player{}

	for i := n; i > 0; i-- {
		for _, p := range cur.getData(year-i, "") {
			if p.ID() == id {
				results = append(results, p)
			}
		}
	}
	return results
}

// GetYear returns a list of all Players, with batting and pitching lines, for the regular season that year.
func GetYear(year int) []*Player { return cur.getData(year, "regular") }

// GetPostYear returns a list of all Players, with batting and pitching lines, for postseason that year.
func GetPostYear(year int) []*Player { return cur.getData(year, "postseason") }

// getData reads battings and pitching data from a given database, merges any duplicate players together,
// and returns the list.
func (d *db) getData(year int, which string) []*Player { /// bat batDB, pit pitDB) []*Player {
	var bat batDB
	var pit pitDB

	d.load(year)

	switch which {
	case "postseason":
		bat, pit = d.batPostDB, d.pitPostDB
	default: // "regular"
		bat, pit = d.batRegularDB, d.pitRegularDB
	}

	batters := battingYear(year, bat)
	pitchers := pitchingYear(year, pit)
	players := make([]*Player, len(batters))

	for i, b := range batters {
		players[i] = &Player{}
		players[i].b = true
		players[i].bio = b.bio
		players[i].master = d.masterDB[b.bio.id]
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

// read the batting database and returns a list of batters for a given year
func battingYear(year int, bat batDB) []*batter {
	res := []*batter{}
	for _, b := range bat {
		if b.year == float64(year) {
			res = append(res, b)
		}
	}
	return res
}

// read the pitching database and return a list of pitchers for a given year
func pitchingYear(year int, pit pitDB) []*pitcher {
	res := []*pitcher{}
	for _, p := range pit {
		if p.year == float64(year) {
			res = append(res, p)
		}
	}
	return res
}

// initializes the pitching and batting databases into package variables
func init() {
	// init the master database
	cur.masterDB = make(map[string]master)
	lines, err := readAll(masterCsv, master{})
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range lines {
		m := l.(master)
		cur.masterDB[m[0]] = m
	}
}

func Load(years ...int) {
	cur.years = []int{} //rest years to force load
	cur.load(years...)
}

func (d *db) checkYears(in []int) (out []int) {
	for _, y := range in {
		already := false
		for _, yC := range cur.years {
			if y == yC {
				already = true
				break
			}
		}
		if !already {
			out = append(out, y)
		}
	}
	return out
}

func (d *db) load(years ...int) {
	newY := d.checkYears(years)

	if len(newY) != 0 {
		//fmt.Println("loading more years", newY)
		d.batPostDB = append(d.batPostDB, initBat(battingPost, newY)...)
		d.pitPostDB = append(d.pitPostDB, initPit(pitchingPost, newY)...)
		d.batRegularDB = append(d.batRegularDB, initBat(batting, newY)...)
		d.pitRegularDB = append(d.pitRegularDB, initPit(pitching, newY)...)
	}

	d.years = append(d.years, newY...)
}

// initializes a batting database
func initBat(file string, years []int) (results []*batter) {
	lines, err := readAll(file, batter{years: years})
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range lines {
		b := l.(*batter)
		results = append(results, b)
	}
	return results
}

// initializes a pitching database
func initPit(file string, years []int) (results []*pitcher) {
	lines, err := readAll(file, pitcher{years: years})
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range lines {
		p := l.(*pitcher)
		results = append(results, p)
	}
	return results
}
