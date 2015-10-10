package lahman

import "fmt"

// A Batter holds all the stats for a player's batting line
type Batter struct {
	player
	G, AB, R, H, H2, H3, HR, RBI, SB, CS, BB, SO, IBB, HBP, SH, SF, GIDP int16
}

/*
func newBatter() *Batter {
	b := &Batter{player: player{}}
	return b
}
*/

// String prints a Batter
func (b Batter) String() string {
	return fmt.Sprintf("%s,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d",
		b.player, b.G, b.AB, b.R, b.H, b.H2, b.H3, b.HR, b.RBI, b.SB, b.CS,
		b.BB, b.SO, b.IBB, b.HBP, b.SH, b.SF, b.GIDP)

}

// csvRead implements csvReader
// It reads from a csv line, and returns an instance of a Batter object.
// Example of use:
//      b, err := Batter{}.csvRead(line)
func (bat Batter) csvRead(line []string) (csvReader, error) {

	b := &Batter{player: player{}}
	ep := &errParser{}

	b.ID = line[0]
	b.Year = ep.parseStat(line[1])
	b.Stint = ep.parseStat(line[2])
	b.Team = line[3]
	b.League = line[4]
	b.G = ep.parseStat(line[5])
	b.AB = ep.parseStat(line[6])
	b.R = ep.parseStat(line[7])
	b.H = ep.parseStat(line[8])
	b.H2 = ep.parseStat(line[9])
	b.H3 = ep.parseStat(line[10])
	b.HR = ep.parseStat(line[11])
	b.RBI = ep.parseStat(line[12])
	b.SB = ep.parseStat(line[13])
	b.CS = ep.parseStat(line[14])
	b.BB = ep.parseStat(line[15])
	b.SO = ep.parseStat(line[16])
	b.IBB = ep.parseStat(line[17])
	b.HBP = ep.parseStat(line[18])
	b.SH = ep.parseStat(line[19])
	b.SF = ep.parseStat(line[20])
	b.GIDP = ep.parseStat(line[21])

	if ep.err != nil {
		return nil, ep.err
	}

	return b, nil
}
