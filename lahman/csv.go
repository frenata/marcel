package lahman

import (
	"bytes"
	"encoding/csv"
	"io/ioutil"
)

// csvReader is an interface that can read a line of csv and create an object
type csvReader interface {
	csvRead([]string) (csvReader, error)
}

// newReader starts up a new csv reader from a file
func newReader(file string) (*csv.Reader, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(bytes.NewReader(b))

	return r, nil
}

// ReadAll reads all the lines from a csv file and creates a list of objects from that file.
func ReadAll(file string) ([]csvReader, error) {
	r, err := newReader(file)
	if err != nil {
		return nil, err
	}

	r.Read() // dispose of first line
	lines, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	results := make([]csvReader, len(lines))
	for i, l := range lines {
		b, err := Batter{}.csvRead(l)
		if err != nil {
			return nil, err
		}
		results[i] = b
	}

	return results, err
}
