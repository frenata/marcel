package lahman

import (
	"fmt"
	"strconv"
)

// BatStats is an array that holds the 18 tracked batting statistics for a Player.
// 17 are imported from the lahmanDB, and PA is calculated from that data.
type BatStats [18]float64

// String returns a list of the 18 battings stats (including PA).
func (b BatStats) String() string {
	var s string

	for i := 1; i < len(b); i++ {
		switch {
		case b[i] == -1:
			s += ","
		default:
			s += fmt.Sprintf("%.0f,", b[i])
		}
	}
	return s[:len(s)-1]
}

// PA, etc, return the Player's named batting stat.
func (b BatStats) PA() float64   { return b[0] }
func (b BatStats) G() float64    { return b[1] }
func (b BatStats) AB() float64   { return b[2] }
func (b BatStats) R() float64    { return b[3] }
func (b BatStats) H() float64    { return b[4] }
func (b BatStats) H2() float64   { return b[5] }
func (b BatStats) H3() float64   { return b[6] }
func (b BatStats) HR() float64   { return b[7] }
func (b BatStats) RBI() float64  { return b[8] }
func (b BatStats) SB() float64   { return b[9] }
func (b BatStats) CS() float64   { return b[10] }
func (b BatStats) BB() float64   { return b[11] }
func (b BatStats) SO() float64   { return b[12] }
func (b BatStats) IBB() float64  { return b[13] }
func (b BatStats) HBP() float64  { return b[14] }
func (b BatStats) SH() float64   { return b[15] }
func (b BatStats) SF() float64   { return b[16] }
func (b BatStats) GIDP() float64 { return b[17] }

// A Batter holds all the stats for a player's batting line
type batter struct {
	bio
	BatStats
}

// String prints a Batter
func (b batter) String() string {
	return fmt.Sprintf("%s,%s",
		b.bio, b.BatStats)

}

// csvRead implements csvReader
// It reads from a csv line, and returns an instance of a Batter object.
// Example of use:
//      b, err := Batter{}.csvRead(line)
func (bat batter) csvRead(line []string) (csvReader, error) {

	b := &batter{bio: bio{}}
	ep := &errParser{}

	_, err := strconv.ParseFloat(line[0], 64)

	switch err {
	case nil: // BattingPost format
		b.id = line[2]
		b.year = ep.parseStat(line[0])
		b.stint = line[1]
		b.team = line[3]
		b.league = line[4]
	default: // all other formats
		b.id = line[0]
		b.year = ep.parseStat(line[1])
		b.stint = line[2]
		b.team = line[3]
		b.league = line[4]
	}

	for i := 1; i < len(b.BatStats); i++ {
		b.BatStats[i] = ep.parseStat(line[i+4])
	}

	b.BatStats[0] = b.BatStats[2] + b.BatStats[11] + b.BatStats[14] + b.BatStats[15] + b.BatStats[16]

	if ep.err != nil {
		return nil, ep.err
	}

	return b, nil
}
