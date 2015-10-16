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
	//fmt.Println(file, c.start(), c.end())
	//start, end := c.start(), c.end()

	/*
			r, err := newReader(file)
			if err != nil {
				return nil, err
			}

			r.Read() // dispose of first line
			lines, err := r.ReadAll()
			if err != nil {
				return nil, err
			}

			results := make([]csvReader, 0) // len(lines))
			fmt.Println(len(lines))
		loop:
			for _, l := range lines {
				year, err := strconv.Atoi(l[1])
				if err != nil { //BAttingPost.csv
					year, err = strconv.Atoi(l[0])
				}
				switch {
				case year < c.start():
					break
				case year > c.end() && c.end() != -1:
					fmt.Println("stop looking")
					break loop
				default:
					b, err := c.csvRead(l)
					if err != nil {
						return nil, err
					}
					results = append(results, b)
				}
			}

			//	fmt.Println(file, len(results))
			return results, err
	*/

	years := []int{}
	for i := c.start(); i < c.end(); i++ {
		years = append(years, i)
	}
	r, err := grepYear(file, years)

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

	return csv.NewReader(stdout), nil
}
