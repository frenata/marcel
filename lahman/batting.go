package lahman

import "fmt"

// A Batter holds all the stats for a player's batting line
type Batter struct {
	player
	BatStats
}

// array of the 17 lahman stats + PA
type BatStats [18]float64

/*
type BatStats struct {
	G, AB, R, H, H2, H3, HR, RBI, SB, CS, BB, SO, IBB, HBP, SH, SF, GIDP float64
}
*/

// String prints a Batter
func (b Batter) String() string {
	return fmt.Sprintf("%s,%s",
		b.player, b.BatStats)

}

func (b BatStats) String() string {
	var s string

	for i := 1; i < len(b); i++ {
		s += fmt.Sprintf("%.0f,", b[i])
	}
	return s[:len(s)-1]
}

/*
func (b BatStats) String() string {
	return fmt.Sprintf("%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f",
		b.G, b.AB, b.R, b.H, b.H2, b.H3, b.HR, b.RBI, b.SB, b.CS,
		b.BB, b.SO, b.IBB, b.HBP, b.SH, b.SF, b.GIDP)
}
*/

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

	for i := 1; i < len(b.BatStats); i++ {
		b.BatStats[i] = ep.parseStat(line[i+4])
	}

	b.BatStats[0] = b.BatStats[2] + b.BatStats[11] + b.BatStats[14] + b.BatStats[15] + b.BatStats[16]

	if ep.err != nil {
		return nil, ep.err
	}

	return b, nil
}

// Convenience methods to return the named stat instead of needing to know the index.
func (b Batter) PA() float64   { return b.BatStats[0] }
func (b Batter) G() float64    { return b.BatStats[1] }
func (b Batter) AB() float64   { return b.BatStats[2] }
func (b Batter) R() float64    { return b.BatStats[3] }
func (b Batter) H() float64    { return b.BatStats[4] }
func (b Batter) H2() float64   { return b.BatStats[5] }
func (b Batter) H3() float64   { return b.BatStats[6] }
func (b Batter) HR() float64   { return b.BatStats[7] }
func (b Batter) RBI() float64  { return b.BatStats[8] }
func (b Batter) SB() float64   { return b.BatStats[9] }
func (b Batter) CS() float64   { return b.BatStats[10] }
func (b Batter) BB() float64   { return b.BatStats[11] }
func (b Batter) SO() float64   { return b.BatStats[12] }
func (b Batter) IBB() float64  { return b.BatStats[13] }
func (b Batter) HBP() float64  { return b.BatStats[14] }
func (b Batter) SH() float64   { return b.BatStats[15] }
func (b Batter) SF() float64   { return b.BatStats[16] }
func (b Batter) GIDP() float64 { return b.BatStats[17] }
