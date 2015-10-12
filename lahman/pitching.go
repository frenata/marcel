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

// String prints a Pitcher
func (p pitcher) String() string {
	return fmt.Sprintf("%s,%s",
		p.bio, p.PitchStats)
}

func (pit pitcher) csvRead(line []string) (csvReader, error) {
	p := &pitcher{bio: bio{}}
	ep := &errParser{}

	p.ID = line[0]
	p.Year = ep.parseStat(line[1])
	p.Stint = line[2] // ep.parseStat(line[2])
	p.Team = line[3]
	p.League = line[4]

	for i := 0; i < len(p.PitchStats); i++ {
		p.PitchStats[i] = ep.parseStat(line[i+5])
	}

	if ep.err != nil {
		//fmt.Println(ep.err)
		return nil, ep.err
	}

	return p, nil
}

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
