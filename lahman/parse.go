package lahman

import "strconv"

// Parse a series of errors
type errParser struct {
	err error
}

// helper function to parse a string into an int16
// returns 0 if data is empty
// stores an error in the errParser struct if there is one
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

func (ep *errParser) parseFloat(s string) float32 {
	if s == "" {
		return 0
	}

	n, err := strconv.ParseFloat(s, 32)

	if err != nil {
		ep.err = err
		return 0
	}

	return float32(n)
}
