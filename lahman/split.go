package lahman

func BattingYear(year int16, batters []*Player) []*Player {
	res := []*Player{}
	for _, b := range batters {
		if b.Year == year {
			res = append(res, b)
		}
	}
	return res
}
