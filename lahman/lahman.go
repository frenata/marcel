package lahman

import (
	"log"
	"os"
)

type batDB []*batter
type pitDB []*pitcher

var (
	master       string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/Master.csv"
	batting      string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/Batting.csv"
	pitching     string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/Pitching.csv"
	battingPost  string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/BattingPost.csv"
	pitchingPost string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/PitchingPost.csv"

	batPostDB    batDB
	pitPostDB    pitDB
	batRegularDB batDB
	pitRegularDB pitDB
	masterDB     map[string]Master
)

// GetYear returns a list of all Players, with batting and pitching lines, for that year.
func GetYear(year float64) []*Player     { return getWhich(year, "regular") }
func GetPostYear(year float64) []*Player { return getWhich(year, "postseason") }

func getWhich(year float64, which string) []*Player {
	switch which {
	case "postseason":
		return getData(year, batPostDB, pitPostDB)
	case "regular":
		return getData(year, batRegularDB, pitRegularDB)
	default:
		return nil
	}
}

func getData(year float64, bat batDB, pit pitDB) []*Player {
	batters := battingYear(year, bat)
	pitchers := pitchingYear(year, pit)
	players := make([]*Player, len(batters))

	for i, b := range batters {
		players[i] = &Player{}
		players[i].b = true
		players[i].bio = b.bio
		players[i].Mas = masterDB[b.bio.ID]
		players[i].bio.First = masterDB[b.bio.ID][13]
		players[i].bio.Last = masterDB[b.bio.ID][14]
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
func battingYear(year float64, bat batDB) []*batter {
	res := []*batter{}
	for _, b := range bat {
		if b.Year == year {
			res = append(res, b)
		}
	}
	return res
}

// read the pitching database and return a list of pitchers for a given year
func pitchingYear(year float64, pit pitDB) []*pitcher {
	res := []*pitcher{}
	for _, p := range pit {
		if p.Year == year {
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
	masterDB = make(map[string]Master)
	mlines, err := ReadAll(master, Master{})
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range mlines {
		m := l.(Master)
		masterDB[m[0]] = m
	}
}

func initBat(file string) (results []*batter) {
	lines, err := ReadAll(file, batter{})
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
	lines, err := ReadAll(file, pitcher{})
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range lines {
		p := l.(*pitcher)
		results = append(results, p)
	}
	return results
}
