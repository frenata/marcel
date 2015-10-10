package lahman

type Player struct {
	ID     string
	Year   uint16
	Stint  uint16
	Team   string
	League string
}

//func (p Player) Year() int  { return int(p.year) }
//func (p Player) Stint() int { return int(p.stint) }
