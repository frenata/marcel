package lahman

import "strconv"

type Batter struct {
	Player
	G, AB, R, H, H2, H3, HR, RBI, SB, CS, BB, SO, IBB, HBP, SH, SF, GIDP uint16
}

func NewBatter(id string, year, stint uint16) (*Batter, error) {
	b := &Batter{Player: Player{ID: id, Year: year, Stint: stint}}
	// search the db for the correct entry, fill in the data

	return b, nil
}

func NewBatterCSV(r []string) (*Batter, error) {
	b := &Batter{Player: Player{}}
	ep := &errParser{}

	b.ID = r[0]
	b.Year = ep.parseStat(r[1])
	b.Stint = ep.parseStat(r[2])
	b.Team = r[3]
	b.League = r[4]
	b.G = ep.parseStat(r[5])
	b.AB = ep.parseStat(r[6])
	b.R = ep.parseStat(r[7])
	b.H = ep.parseStat(r[8])
	b.H2 = ep.parseStat(r[9])
	b.H3 = ep.parseStat(r[10])
	b.HR = ep.parseStat(r[11])
	b.RBI = ep.parseStat(r[12])
	b.SB = ep.parseStat(r[13])
	b.CS = ep.parseStat(r[14])
	b.BB = ep.parseStat(r[15])
	b.SO = ep.parseStat(r[16])
	b.IBB = ep.parseStat(r[17])
	b.HBP = ep.parseStat(r[18])
	b.SH = ep.parseStat(r[19])
	b.SF = ep.parseStat(r[20])
	b.GIDP = ep.parseStat(r[21])

	if ep.err != nil {
		return nil, ep.err
	}

	return b, nil
}

func (ep *errParser) parseStat(s string) uint16 {
	if s == "" {
		return 0
	}

	n, err := strconv.ParseUint(s, 10, 16)

	if err != nil {
		ep.err = err
		return 0
	}

	return uint16(n)
}

type errParser struct {
	err error
}
