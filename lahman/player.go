package lahman

import (
	"fmt"
	"strconv"
)

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
func (p Player) IsPitcher() bool { return p.p }
func (p Player) IsBatter() bool  { return p.b }

// String returns a string representing a Player's complete statline.
func (p Player) String() string {
	var s string
	if name := p.Name(); name != "" {
		s = name + ","
	}
	s += p.bio.String()
	if p.b {
		s += "\nBatting:  " + p.Bat.String()
	}
	if p.p {
		s += "\nPitching: " + p.Pit.String()
	}
	return s
}

// struct that list the unique identifying information for a given player's stats
// for embedding into more complete structs: Batter, Pitcher, Fielder
type bio struct {
	id     string
	year   float64
	stint  string
	team   string
	league string
}

// String returns the basic information about a Player's bio
func (p bio) String() string {
	return fmt.Sprintf("%s,%.0f,%s,%s,%s", p.id, p.year, p.stint, p.team, p.league)
}

// ID returns a player's lahmanDB ID
func (p bio) ID() string { return p.id }

// Year returns the year of a Player's statline.
func (p bio) Year() int { return int(p.year) }

// Stint returns:
// 	for regular season play, the number of their season's appearance with a given team
//	(In other words, if a players is traded their stats for the second team will be stint "2")
//	for postseason play, the *round* of play
func (p bio) Stint() string { return p.stint }

// Team returns the team code a Player played for.
func (p bio) Team() string { return p.team }

// League returns the league a Player played for.
func (p bio) League() string { return p.league }

// Master lists biographical data fromt he Master database
type master [24]string

func (m master) yearS() []int { return []int{} }

//csvRead pulls data from a csv file into a Master object
func (mas master) csvRead(line []string) (csvReader, error) {
	m := master{}

	for i := 0; i < len(line); i++ {
		m[i] = line[i]
	}

	return m, nil
}

// Name returns a player's name.
func (p master) Name() string { return fmt.Sprintf("%s %s", p[13], p[14]) }

// Birth returns a player's birth year.
func (p master) Birth() int {
	y := p[1]
	n, _ := strconv.Atoi(y)
	return n
}
