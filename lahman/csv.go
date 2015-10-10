package lahman

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
)

func ReadCSV(file string) (*csv.Reader, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	//	buff := bytes.NewBuffer(b)
	//	s := b.String()

	r := csv.NewReader(bytes.NewReader(b))

	return r, nil
}

func ReadAll(file string) ([]*Batter, error) {
	r, err := ReadCSV(file)
	if err != nil {
		return nil, err
	}

	r.Read() // dispose of first line
	lines, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	batters := make([]*Batter, len(lines))
	for i, l := range lines {
		//err := batters[i].csvRead(l)
		b, err := NewBatterCSV(l)
		if err != nil {
			return nil, err
		}
		batters[i] = b
	}

	return batters, err
}
