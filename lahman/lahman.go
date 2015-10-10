package lahman

import "errors"

// BattingYear returns all the records from a given year from a list of Batters.
func BattingYear(year int16, lines []csvReader) ([]*Batter, error) {
	res := []*Batter{}
	for _, l := range lines {
		b, ok := l.(*Batter)
		if !ok {
			return nil, errors.New("Not a list of Batting lines")
		}
		if b.Year == year {
			res = append(res, b)
		}
	}
	return res, nil
}

func PitchingYear(year int16, lines []csvReader) ([]*Pitcher, error) {
	res := []*Pitcher{}
	for _, l := range lines {
		p, ok := l.(*Pitcher)
		if !ok {
			return nil, errors.New("Not a list of Pitching lines")
		}
		if p.Year == year {
			res = append(res, p)
		}
	}
	return res, nil
}
