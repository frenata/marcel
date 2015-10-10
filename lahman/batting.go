package lahman

import (
	"fmt"
	"strconv"
)

type Batter struct {
	Player
	G, AB, R, H, H2, H3, HR, RBI, SB, CS, BB, SO, IBB, HBP, SH, SF, GIDP int16
}

func (b Batter) String() string {
	return fmt.Sprintf("%s,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d",
		b.Player, b.G, b.AB, b.R, b.H, b.H2, b.H3, b.HR, b.RBI, b.SB, b.CS,
		b.BB, b.SO, b.IBB, b.HBP, b.SH, b.SF, b.GIDP)

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

func (ep *errParser) parseStat(s string) int16 {
	if s == "" {
		return 0
	}

	n, err := strconv.ParseInt(s, 10, 16)

	if err != nil {
		ep.err = err
		return 0
	}

	return int16(n)
}

type errParser struct {
	err error
}
