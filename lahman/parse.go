package lahman

import "strconv"

// Parse a series of errors
type errParser struct {
	err error
}

func (ep *errParser) parseStat(s string) float64 {
	if s == "" {
		return -1
	}

	n, err := strconv.ParseFloat(s, 64)

	if err != nil {
		ep.err = err
		return 0
	}

	return n
}
