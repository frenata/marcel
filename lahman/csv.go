package lahman

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strconv"
)

// csvReader is an interface that can read a line of csv and create an object
type csvReader interface {
	csvRead([]string) (csvReader, error)
	start() int
	end() int
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
func readAll(file string, c csvReader) ([]csvReader, error) {
	//fmt.Println(file, c.start(), c.end())
	//start, end := c.start(), c.end()

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
loop:
	for _, l := range lines {
		year, _ := strconv.Atoi(l[1])
		switch {
		case year < c.start():
			break
		case year > c.end() && c.end() != -1:
			break loop
		default:
			b, err := c.csvRead(l)
			if err != nil {
				return nil, err
			}
			results = append(results, b)
		}
	}

	fmt.Println(len(results))
	return results, err
}
