package marcel

import (
	"fmt"

	"github.com/frenata/marcel/lahman"
)

type Player struct {
	lahman.Player
}

func regressPlayer(id string, year int) (regress lahman.BatStats) {
	accum, pa, birth := weightPlayer(id, year)
	league1, _ := leagueAvg(year - 1)
	league2, _ := leagueAvg(year - 2)
	league3, _ := leagueAvg(year - 3)

	playtime := pa[0]*.5 + pa[1]*.1 + 200
	//fmt.Println("playtime", playtime)

	for i := 1; i < len(accum); i++ {
		L := (league1[i]*pa[0]*5 +
			league2[i]*pa[1]*4 +
			league3[i]*pa[2]*3) /
			accum.PA() * 1200
		regress[i] = (accum[i] + L) / (accum.PA() + 1200) * playtime
	}

	var age, ageAdj float64
	age = float64(year) - float64(birth)
	switch {
	case age > 29:
		ageAdj = (age - 29) * .003
	case age < 29:
		ageAdj = (29 - age) * .006
	default:
		ageAdj = 0
	}

	fmt.Println(ageAdj)
	for i := 3; i < len(regress); i++ {
		switch {
		case i == 10:
			fallthrough
		case i == 12:
			fallthrough
		case i == 17:
			regress[i] = regress[i] * (1 - ageAdj)
		default: // good events
			regress[i] = regress[i] * (1 + ageAdj)
		}
	}
	regress[0] = playtime

	//age = 2004 - yearofbirth. If over 29, AgeAdj = (age - 29) * .003. If under 29, AgeAdj = (age - 29) * .006. Apply this age adjustment to the result of #4.
	return regress
}

func weightPlayer(id string, year int) (accum lahman.BatStats, pa [3]float64, birth int) {
	//accum := lahman.BatStats{}
	res := lahman.GetPlayer(id, year, 3)
	for _, p := range res {
		for i, v := range p.Bat {
			switch {
			case p.Year() == year-1:
				pa[0] = p.Bat.PA()
				accum[i] = accum[i] + v*5
			case p.Year() == year-2:
				pa[1] = p.Bat.PA()
				accum[i] = accum[i] + v*4
			case p.Year() == year-3:
				pa[2] = p.Bat.PA()
				accum[i] = accum[i] + v*3
			}
		}
	}
	return accum, pa, res[0].Birth()
}

func leagueAvg(year int) (bat lahman.BatStats, pit lahman.PitchStats) {
	var bCount, pCount, ipCount, bfCount float64
	players := lahman.GetYear(year)

	for _, p := range players {
		if p.IsBatter() {
			for i := 0; i < len(bat); i++ {
				bat[i] += p.Bat[i]
			}
			bCount++
		}
		if p.IsPitcher() && !p.IsPosPitcher() {
			for i := 0; i < len(pit); i++ {
				switch i {
				case 13:
					pit[i] = pit[i] + p.Pit[i]*p.Pit.BFP()
					bfCount += p.Pit.BFP()
				case 14:
					pit[i] = pit[i] + p.Pit[i]*p.Pit.IPouts()
					ipCount += p.Pit.IPouts()
				default:
					pit[i] += p.Pit[i]
				}
			}
			pCount++
		}
	}

	for i := 1; i < len(bat); i++ {
		bat[i] = bat[i] / bat.PA()
	}
	bat[0] = bat[0] / bCount

	for i := 0; i < len(pit); i++ {
		switch i {
		case 13:
			pit[i] = pit[i] / bfCount
		case 14:
			pit[i] = pit[i] / ipCount
		default:
			pit[i] = pit[i] / pCount
		}
	}
	return bat, pit
}
