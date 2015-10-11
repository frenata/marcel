package lahman

import "fmt"

// A Pitcher holds all the stats for a player's pitching line
type Pitcher struct {
	player
	PitchStats
}

type PitchStats struct {
	W, L, G, GS, CG, SHO, SV, IPouts, H, ER, HR, BB,
	SO, IBB, WP, HBP, BK, BFP, GF, R, SH, SF, GIDP,
	BAopp, ERA float64
}

// String prints a Pitcher
func (p Pitcher) String() string {
	return fmt.Sprintf("%s,%s",
		p.player, p.PitchStats)
}

func (p PitchStats) String() string {
	return fmt.Sprintf(
		"%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,%.0f,"+
			"%.0f,%.0f,%.0f,%4.3f,%3.2f,%.0f,%.0f,%.0f,%.0f,%.0f,"+
			"%.0f,%.0f,%.0f,%.0f,%.0f",
		p.W, p.L, p.G, p.GS, p.CG, p.SHO, p.SV, p.IPouts,
		p.H, p.ER, p.HR, p.BB, p.SO, p.BAopp, p.ERA, p.IBB, p.WP,
		p.HBP, p.BK, p.BFP, p.GF, p.R, p.SH, p.SF, p.GIDP)
}

// csvRead implements csvReader
// It reads from a csv line, and returns an instance of a Pitcher object.
// Example of use:
//      p, err := Pitcher{}.csvRead(line)
func (pit Pitcher) csvRead(line []string) (csvReader, error) {

	p := &Pitcher{player: player{}}
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
