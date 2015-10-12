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

// GetYear returns a list of all Players, with batting and pitching lines, for that year.
func GetYear(year int) []*Player     { return getWhich(year, "regular") }
func GetPostYear(year int) []*Player { return getWhich(year, "postseason") }

func getWhich(year int, which string) []*Player {
	switch which {
	case "postseason":
		return getData(year, batPostDB, pitPostDB)
	case "regular":
		return getData(year, batRegularDB, pitRegularDB)
	default:
		return nil
	}
}

func getData(year int, bat batDB, pit pitDB) []*Player {
	batters := battingYear(year, bat)
	pitchers := pitchingYear(year, pit)
	players := make([]*Player, len(batters))

	for i, b := range batters {
		players[i] = &Player{}
		players[i].b = true
		players[i].bio = b.bio
		players[i].master = masterDB[b.bio.ID]
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
		if b.Year == float64(year) {
			res = append(res, b)
		}
	}
	return res
}

// read the pitching database and return a list of pitchers for a given year
func pitchingYear(year int, pit pitDB) []*pitcher {
	res := []*pitcher{}
	for _, p := range pit {
		if p.Year == float64(year) {
			res = append(res, p)
		}
	}
	return res
}

// initializes the pitching and batting databases into package variables
func init() {
	batRegularDB = initBat(batting)
	pitRegularDB = initPit(pitching)

	batPostDB = initBat(battingPost)
	pitPostDB = initPit(pitchingPost)

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

func initBat(file string) (results []*batter) {
	lines, err := readAll(file, batter{})
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range lines {
		b := l.(*batter)
		results = append(results, b)
	}
	return results
}

func initPit(file string) (results []*pitcher) {
	lines, err := readAll(file, pitcher{})
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range lines {
		p := l.(*pitcher)
		results = append(results, p)
	}
	return results
}
