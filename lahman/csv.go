package lahman

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os/exec"
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
	var r *csv.Reader
	var err error

	if c.start() == 0 { // no limits on db load
		r, err = newReader(file)
		if err != nil {
			return nil, err
		}
		r.Read() //dispose of the first line if not grepping file
	} else { // limited db load
		years := []int{}
		for i := c.start(); i < c.end()+1; i++ {
			years = append(years, i)
		}
		r, err = grepYear(file, years)
		if err != nil {
			return nil, err
		}
	}

	lines, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	results := make([]csvReader, len(lines))
	for i, l := range lines {
		b, err := c.csvRead(l)
		if err != nil {
			return nil, err
		}
		results[i] = b
	}

	return results, err
}

func grepYear(file string, years []int) (*csv.Reader, error) {
	var year string
	for _, y := range years {
		year = fmt.Sprintf("%s\\|%s", year, strconv.Itoa(y))
	}
	year = year[2:]

	cmd := exec.Command("grep", year, file)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(stdout)

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	//fmt.Println(len(buf))
	return csv.NewReader(bytes.NewReader(buf)), nil
}
