package lahman

import (
	"log"
	"os"
)

type batDB []*batter
type pitDB []*pitcher

var (
	masterCsv    string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/Master.csv"
	batting      string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/Batting.csv"
	pitching     string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/Pitching.csv"
	battingPost  string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/BattingPost.csv"
	pitchingPost string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/PitchingPost.csv"

	batPostDB    batDB
	pitPostDB    pitDB
	batRegularDB batDB
	pitRegularDB pitDB
	masterDB     map[string]master
)

// GetPlayer returns the past 3 years of data for that Player
// TODO: test
func GetPlayer(id string, year int, n int) []*Player {
	results := []*Player{}

	for i := n; i > 0; i-- {
		for _, p := range getWhich(year-i, "") {
			if p.ID() == id {
				results = append(results, p)
			}
		}
	}
	return results
}

// GetYear returns a list of all Players, with batting and pitching lines, for the regular season that year.
func GetYear(year int) []*Player { return getWhich(year, "regular") }

// GetPostYear returns a list of all Players, with batting and pitching lines, for postseason that year.
func GetPostYear(year int) []*Player { return getWhich(year, "postseason") }

// getWhich picks what databases to read based on the helper function that calls it.
func getWhich(year int, which string) []*Player {
	switch which {
	case "postseason":
		return getData(year, batPostDB, pitPostDB)
	case "regular":
		fallthrough
	default:
		return getData(year, batRegularDB, pitRegularDB)
	}
}

// getData reads battings and pitching data from a given database, merges any duplicate players together,
// and returns the list.
func getData(year int, bat batDB, pit pitDB) []*Player {
	batters := battingYear(year, bat)
	pitchers := pitchingYear(year, pit)
	players := make([]*Player, len(batters))

	for i, b := range batters {
		players[i] = &Player{}
		players[i].b = true
		players[i].bio = b.bio
		players[i].master = masterDB[b.bio.id]
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
	batRegularDB = initBat(batting, 2000, 2015)
	pitRegularDB = initPit(pitching, 2000, 2015)

	batPostDB = initBat(battingPost, 2000, 2015)
	pitPostDB = initPit(pitchingPost, 2000, 2015)

	// init the master database
	masterDB = make(map[string]master)
	lines, err := readAll(masterCsv, master{})
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range lines {
		m := l.(master)
		masterDB[m[0]] = m
	}
}

// initializes a batting database
func initBat(file string, start, end int) (results []*batter) {
	lines, err := readAll(file, batter{startY: start, endY: end})
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
func initPit(file string, start, end int) (results []*pitcher) {
	lines, err := readAll(file, pitcher{startY: start, endY: end})
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range lines {
		p := l.(*pitcher)
		results = append(results, p)
	}
	return results
}
