package lahman

type Player struct {
	ID     string
	Year   int16
	Stint  int16
	Team   string
	League string
}

//func (p Player) Year() int  { return int(p.year) }
//func (p Player) Stint() int { return int(p.stint) }
