package lahman

import "os"

var bats string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/Batting.csv"
var pits string = os.Getenv("GOPATH") + "/src/github.com/frenata/marcel/lahman/data/Pitching.csv"

var BatDB []*Batter
var PitDB []*Pitcher

func init() {
	lines, _ := ReadAll(bats, Batter{})
	res := []*Batter{}
	for _, l := range lines {
		b := l.(*Batter)
		res = append(res, b)
	}
	BatDB = res

	plines, _ := ReadAll(pits, Pitcher{})
	pres := []*Pitcher{}
	for _, l := range plines {
		b := l.(*Pitcher)
		pres = append(pres, b)
	}
	PitDB = pres
}

func BattingYear(year int16) []*Batter {
	res := []*Batter{}
	for _, b := range BatDB {
		if b.Year == year {
			res = append(res, b)
		}
	}
	return res
}

func PitchingYear(year int16) []*Pitcher {
	res := []*Pitcher{}
	for _, p := range PitDB {
		if p.Year == year {
			res = append(res, p)
		}
	}
	return res
}

/*
// BattingYear returns all the records from a given year from a list of Batters.
func BattingYear(year int16, file string) ([]*Batter, error) {
	lines, err := ReadAll(file, Batter{})
	if err != nil {
		return nil, err
	}
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

func PitchingYear(year int16, file string) ([]*Pitcher, error) {
	lines, err := ReadAll(file, Pitcher{})
	if err != nil {
		return nil, err
	}
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
*/
