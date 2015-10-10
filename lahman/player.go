package lahman

type Player struct {
	ID     string
	Year   int16
	Stint  int16
	Team   string
	League string
	Batter
}

func NewPlayer(id string, year, stint int16) (*Player, error) {
	p := &Player{ID: id, Year: year, Stint: stint}
	// search the db for the correct entry, fill in the data

	return p, nil
}

//func (p Player) Year() int  { return int(p.year) }
//func (p Player) Stint() int { return int(p.stint) }
