package lahman

import "fmt"

// struct that list the unique identifying information for a given player's stats
// for embedding into more complete structs: Batter, Pitcher, Fielder
type player struct {
	ID     string
	Year   int16
	Stint  int16
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
func (p player) String() string {
	return fmt.Sprintf("%s,%d,%d,%s,%s",
		p.ID, p.Year, p.Stint, p.Team, p.League)

}
