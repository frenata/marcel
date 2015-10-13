package marcel

import "github.com/frenata/marcel/lahman"

type Player struct {
	lahman.Player
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
		if p.IsPitcher() {
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

	for i := 0; i < len(bat); i++ {
		bat[i] = bat[i] / bCount
	}
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
