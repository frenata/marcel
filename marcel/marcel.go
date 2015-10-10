package marcel

import (
	"fmt"
	"log"
	"os"

	"github.com/frenata/marcel/lahman"
)

var bats string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/Batting.csv"
var pits string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/Pitching.csv"

type Player struct {
	Batting  *lahman.Batter
	Pitching *lahman.Pitcher
}

func (p Player) String() string {
	if p.Pitching != nil {
		return fmt.Sprintf("%s\n%s", p.Batting, p.Pitching)
	}
	return fmt.Sprint(p.Batting)
}

func GetPitching(year int16) []*Player {
	pityear, err := lahman.PitchingYear(year, pits)
	if err != nil {
		log.Fatal("wrong list of objects")
	}

	p := make([]*Player, len(pityear))

	for i, py := range pityear {
		p[i] = &Player{}
		p[i].Pitching = py
	}
	return p
}

func GetBatting(year int16, players []*Player) []*Player {
	batyear, err := lahman.BattingYear(year, bats)
	if err != nil {
		log.Fatal("wrong list of objects")
	}

	for _, by := range batyear {
		for _, p := range players {
			if p.Pitching != nil &&
				by.ID == p.Pitching.ID &&
				by.Stint == p.Pitching.Stint &&
				by.Year == p.Pitching.Year {
				p.Batting = by
				break
			}
		}
	}
	return players
}

func GetYear(year int16) []*Player {
	p := GetPitching(year)
	p = GetBatting(year, p)

	return p
}
