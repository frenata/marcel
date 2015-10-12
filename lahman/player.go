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

// String prints biographical information about a player.
func (p bio) String() string {
	return fmt.Sprintf("%s,%.0f,%s,%s,%s", p.ID, p.Year, p.Stint, p.Team, p.League)
}

// A Player is a struct that contains all batting and pitching stats for a player.
type Player struct {
	Bat BatStats
	Pit PitchStats
	bio
	master
	p, b bool // p is true if player pitched, b is true if player batted and did *not* pitch
}

// IsPosPitcher returns true if the player was normally a position player, but pitched at least once.
func (p Player) IsPosPitcher() bool {
	return p.b && p.p
}

// String prints a player's information
func (p Player) String() string {
	var name string
	if f, l := p.FirstName(), p.LastName(); f != "" && l != "" {
		name = fmt.Sprintf("%s %s,", f, l)
	}
	s := name + p.bio.String()
	if p.b {
		s += "\nBatting:  " + p.Bat.String()
	}
	if p.p {
		s += "\nPitching: " + p.Pit.String()
	}
	return s
}

// Master lists biographical data fromt he Master database
type master [24]string

//csvRead pulls data from a csv file into a Master object
func (mas master) csvRead(line []string) (csvReader, error) {
	m := master{}

	for i := 0; i < len(line); i++ {
		m[i] = line[i]
	}

	return m, nil
}

func (p master) FirstName() string { return m[13] }
func (p master) LastName() string  { return m[14] }
