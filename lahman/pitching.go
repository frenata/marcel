package lahman

import "fmt"

// A Pitcher holds all the stats for a player's pitching line
type pitcher struct {
	bio
	PitchStats
}

type PitchStats [25]float64

func (p PitchStats) String() string {
	var s string

	for i := 0; i < len(p); i++ {
		var s2 string
		switch i {
		case 13: // BAopp
			s2 = fmt.Sprintf("%4.3f,", p[i])
		case 14: // ERA
			s2 = fmt.Sprintf("%3.2f,", p[i])
		default:
			s2 = fmt.Sprintf("%.0f,", p[i])
		}
		s += s2
	}
	return s[:len(s)-1]
}

/*
type PitchStats struct {
	W, L, G, GS, CG, SHO, SV, IPouts, H, ER, HR, BB,
	SO, IBB, WP, HBP, BK, BFP, GF, R, SH, SF, GIDP,
	BAopp, ERA float64
}
*/

// String prints a Pitcher
func (p pitcher) String() string {
	return fmt.Sprintf("%s,%s",
		p.bio, p.PitchStats)
}

/*
func (p PitchStats) String() string {
	return fmt.Sprintf(
		"%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,"+
			"%.0f,%.0f,%.0f,%4.3f,%3.2f,%.0f,%.0f,%.0f,%.0f,%.0f,"+
			"%.0f,%.0f,%.0f,%.0f,%.0f",
		p.W, p.L, p.G, p.GS, p.CG, p.SHO, p.SV, p.IPouts,
		p.H, p.ER, p.HR, p.BB, p.SO, p.BAopp, p.ERA, p.IBB, p.WP,
		p.HBP, p.BK, p.BFP, p.GF, p.R, p.SH, p.SF, p.GIDP)
}
*/

func (pit pitcher) csvRead(line []string) (csvReader, error) {
	p := &pitcher{bio: bio{}}
	ep := &errParser{}

	p.ID = line[0]
	p.Year = ep.parseStat(line[1])
	p.Stint = ep.parseStat(line[2])
	p.Team = line[3]
	p.League = line[4]

	for i := 0; i < len(p.PitchStats); i++ {
		p.PitchStats[i] = ep.parseStat(line[i+5])
	}

	//b.BatStats[0] = b.BatStats[2] + b.BatStats[11] + b.BatStats[14] + b.BatStats[15] + b.BatStats[16]

	if ep.err != nil {
		return nil, ep.err
	}

	return p, nil
}

/*
// csvRead implements csvReader
// It reads from a csv line, and returns an instance of a Pitcher object.
// Example of use:
//      p, err := Pitcher{}.csvRead(line)
func (pit pitcher) csvRead(line []string) (csvReader, error) {

	p := &pitcher{bio: bio{}}
	ep := &errParser{}

	p.ID = line[0]
	p.Year = ep.parseStat(line[1])
	p.Stint = ep.parseStat(line[2])
	p.Team = line[3]
	p.League = line[4]
	p.W = ep.parseStat(line[5])
	p.L = ep.parseStat(line[6])
	p.G = ep.parseStat(line[7])
	p.GS = ep.parseStat(line[8])
	p.CG = ep.parseStat(line[9])
	p.SHO = ep.parseStat(line[10])
	p.SV = ep.parseStat(line[11])
	p.IPouts = ep.parseStat(line[12])
	p.H = ep.parseStat(line[13])
	p.ER = ep.parseStat(line[14])
	p.HR = ep.parseStat(line[15])
	p.BB = ep.parseStat(line[16])
	p.SO = ep.parseStat(line[17])
	p.BAopp = ep.parseStat(line[18]) // float
	p.ERA = ep.parseStat(line[19])   // float
	p.IBB = ep.parseStat(line[20])
	p.WP = ep.parseStat(line[21])
	p.HBP = ep.parseStat(line[22])
	p.BK = ep.parseStat(line[23])
	p.BFP = ep.parseStat(line[24])
	p.GF = ep.parseStat(line[25])
	p.R = ep.parseStat(line[26])
	p.SH = ep.parseStat(line[27])
	p.SF = ep.parseStat(line[28])
	p.GIDP = ep.parseStat(line[29])

	if ep.err != nil {
		return nil, ep.err
	}

	return p, nil
}
*/

func (p PitchStats) W() float64      { return p[0] }
func (p PitchStats) L() float64      { return p[1] }
func (p PitchStats) G() float64      { return p[2] }
func (p PitchStats) GS() float64     { return p[3] }
func (p PitchStats) CG() float64     { return p[4] }
func (p PitchStats) SHO() float64    { return p[5] }
func (p PitchStats) SV() float64     { return p[6] }
func (p PitchStats) IPouts() float64 { return p[7] }
func (p PitchStats) H() float64      { return p[8] }
func (p PitchStats) ER() float64     { return p[9] }
func (p PitchStats) HR() float64     { return p[10] }
func (p PitchStats) BB() float64     { return p[11] }
func (p PitchStats) SO() float64     { return p[12] }
func (p PitchStats) BAopp() float64  { return p[13] }
func (p PitchStats) ERA() float64    { return p[14] }
func (p PitchStats) IBB() float64    { return p[15] }
func (p PitchStats) WP() float64     { return p[16] }
func (p PitchStats) HBP() float64    { return p[17] }
func (p PitchStats) BK() float64     { return p[18] }
func (p PitchStats) BFP() float64    { return p[19] }
func (p PitchStats) GF() float64     { return p[20] }
func (p PitchStats) R() float64      { return p[21] }
func (p PitchStats) SH() float64     { return p[22] }
func (p PitchStats) SF() float64     { return p[23] }
func (p PitchStats) GIDP() float64   { return p[24] }
