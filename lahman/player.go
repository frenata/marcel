package lahman

import "fmt"

type Player struct {
	ID     string
	Year   int16
	Stint  int16
	Team   string
	League string
}

func NewPlayer(id string, year, stint int16) (*Player, error) {
	p := &Player{ID: id, Year: year, Stint: stint}

	return p, nil
}

func (p Player) String() string {
	return fmt.Sprintf("%s,%d,%d,%s,%s",
		p.ID, p.Year, p.Stint, p.Team, p.League)

}

//func (p Player) Year() int  { return int(p.year) }
//func (p Player) Stint() int { return int(p.stint) }
