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
	First  string
	Last   string
}

// String prints biographical information about a player.
func (p bio) String() string {
	var name string
	if p.First != "" && p.Last != "" {
		name = fmt.Sprintf("%s %s,", p.First, p.Last)
	}
	basics := fmt.Sprintf("%s,%.0f,%s,%s,%s", p.ID, p.Year, p.Stint, p.Team, p.League)
	return name + basics
}

// A Player is a struct that contains all batting and pitching stats for a player.
type Player struct {
	bio
	Bat BatStats
	Pit PitchStats
	Mas Master

	p, b bool // p is true if player pitched, b is true if player batted and did *not* pitch
}

// IsPosPitcher returns true if the player was normally a position player, but pitched at least once.
func (p Player) IsPosPitcher() bool {
	return p.b && p.p
}

// String prints a player's information
func (p Player) String() string {
	s := p.bio.String()
	if p.b {
		s += "\nBatting:  " + p.Bat.String()
	}
	if p.p {
		s += "\nPitching: " + p.Pit.String()
	}
	return "\n" + s
}

// Master lists biographical data fromt he Master database
type Master [24]string

//csvRead pulls data from a csv file into a Master object
func (mas Master) csvRead(line []string) (csvReader, error) {
	m := Master{}

	for i := 0; i < len(line); i++ {
		m[i] = line[i]
	}

	return m, nil
}
