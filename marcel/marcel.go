package marcel

import (
	"fmt"
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

func GetBatting(year float64) []*Player {
	batyear := lahman.BattingYear(year)

	p := make([]*Player, len(batyear))

	for i, by := range batyear {
		p[i] = &Player{}
		p[i].Batting = by
	}
	return p
}

func GetPitching(year float64, players []*Player) []*Player {
	pityear := lahman.PitchingYear(year)

	for _, py := range pityear {
		for _, p := range players {
			if py.ID == p.Batting.ID &&
				py.Stint == p.Batting.Stint &&
				py.Year == p.Batting.Year {
				p.Pitching = py
				break
			}
		}
	}
	return players
}

func GetYear(year float64) []*Player {
	p := GetBatting(year)
	p = GetPitching(year, p)

	return p
}

func IsPosPitcher(p *Player) bool {
	if p.Pitching != nil && p.Pitching.BFP < p.Batting.AB {
		return true
	}
	return false
}
