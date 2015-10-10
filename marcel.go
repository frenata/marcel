package main

import (
	"fmt"
	"log"

	"github.com/frenata/marcel/lahman"
)

const bats string = "lahman/data/Batting.csv"
const pits string = "lahman/data/Pitching.csv"

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

func GetBats(year int16) []*Player {
	batyear, err := lahman.BattingYear(year, bats)
	if err != nil {
		log.Fatal("wrong list of objects")
	}

	p := make([]*Player, len(batyear))

	for i, by := range batyear {
		p[i] = &Player{}
		p[i].Batting = by
		//_, _ = i, by
	}
	//fmt.Println(p)
	return p
}

func GetPitching(year int16, players []*Player) []*Player {
	pityear, err := lahman.PitchingYear(year, pits)
	if err != nil {
		log.Fatal("wrong list of objects")
	}

	for _, py := range pityear {
		for _, p := range players {
			if py.ID == p.Batting.ID && py.Stint == p.Batting.Stint && py.Year == p.Batting.Year {
				p.Pitching = py
				break
			}
		}
	}
	return players
}

func GetYear(year int16) []*Player {
	p := GetBats(year)
	p = GetPitching(year, p)

	return p
}
