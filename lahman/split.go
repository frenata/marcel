package lahman

func BattingYear(year int16, batters []*Batter) []*Batter {
	res := []*Batter{}
	for _, b := range batters {
		if b.Year == year {
			res = append(res, b)
		}
	}
	return res
}
