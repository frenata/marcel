package lahman

import "fmt"

// struct that list the unique identifying information for a given player's stats
// for embedding into more complete structs: Batter, Pitcher, Fielder
type bio struct {
	ID     string
	Year   float64
	Stint  string
	Team   string
	League string
}

/*
func newPlayer(id string, year, stint int16) (*Player, error) {
	p := &Player{ID: id, Year: year, Stint: stint}

	return p, nil
}
*/

// print a player
func (p bio) String() string {
	return fmt.Sprintf("%s,%.0f,%s,%s,%s",
		p.ID, p.Year, p.Stint, p.Team, p.League)

}

type Player struct {
	bio
	Bat BatStats
	Pit PitchStats

	p, b bool // p is true if player pitched, b is true if player batted and did *not* pitch
}

// IsPosPitcher returns true if the player was normally a position player, but pitched at least once.
func (p Player) IsPosPitcher() bool {
	return p.b && p.p
}

func (p Player) String() string {
	s := p.bio.String()
	if p.b {
		s += "\nBatting:  " + p.Bat.String()
	}
	if p.p {
		s += "\nPitching: " + p.Pit.String()
	}
	return s
}
