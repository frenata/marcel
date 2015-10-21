package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/frenata/marcel/lahman"
)

func Query(s string) (r string, err error) {
	queries := strings.Split(s, ",")
	if len(queries) <= 1 {
		return "", errors.New("No Query Found")
	}

	players, err := valYears(queries[0])
	if err != nil {
		return "", err
	}
	//players := playersQ
	//fmt.Println(queries, len(queries))

	var bat, pit bool
	res := playerS{}
	for i := 1; i < len(queries); i++ {
		q := strings.TrimSpace(queries[i])
		if len(q) == 0 {
			return "", errors.New("empty query")
		}
		if string(q[0]) == "b" {
			bat = true
		}
		if string(q[0]) == "p" {
			pit = true
		}
		res, err = query(q, players)
		if err != nil {
			return "", err
		}
		//results = append(results, res...)
		players = res
	}

	if len(res) == 0 {
		return "No matches found.", nil
	}

	return res.SmartString(bat, pit), nil
}

func query(s string, players playerS) (playerS, error) {
	results := playerS{}
	bat, stat, op, n, err := validate(s)
	if err != nil {
		return players, err
	}

	for _, p := range players {
		//if bmap[query[1]](p.Bat) >= 70 { // needs to be abstracted
		switch bat {
		case true:
			bstat := stat.(func(lahman.BatStats) float64)
			if op(bstat(p.Bat), n) {
				results = append(results, p)
			}
		case false:
			pstat := stat.(func(lahman.PitchStats) float64)
			if op(pstat(p.Pit), n) {
				results = append(results, p)
			}
		}
	}

	if len(players) == 0 {
		return players, errors.New("No results found")
	}
	return results, nil
}

func validate(s string) (bat bool, stat interface{}, op func(x, y float64) bool, n float64, err error) {
	query := strings.Split(s, " ")
	if len(query) != 4 {
		err = errors.New("query must consist of 4 items")
		return
	}
	//fmt.Println(query)

	var statmap map[string]interface{}
	switch query[0] {
	case "b":
		bat = true
		statmap = bmap
	case "p":
		bat = false
		statmap = pmap
	default:
		err = errors.New("first argument must be either 'b' or 'p' to designate batting or pitching stats")
		return
	}

	//stat = stat.(statType)
	stat, ok := statmap[query[1]]
	if !ok {
		err = errors.New("stat " + query[1] + " not recognized")
		return
	}

	switch query[2] {
	case ">":
		op = gThan
	case "<":
		op = lThan
	case ">=":
		op = geThan
	case "<=":
		op = leThan
	case "=":
		fallthrough
	case "==":
		op = equal
	default:
		err = errors.New("operator " + query[2] + " not recognized")
		return
	}

	n, err = strconv.ParseFloat(query[3], 64)
	if err != nil {
		return
	}

	return bat, stat, op, n, nil
}

func valYears(s string) (playerS, error) {
	res := playersQ // start point

	switch {
	case s == "*":
		return res, nil
	case s == "m": // modern era
		return getyears(intslice(FirstModernYear, LastYear)), nil
	case s == "pm": // pre-modern era
		return getyears(intslice(FirstYear, FirstModernYear-1)), nil
	case len(s) < 9:
		break
	default: // 18xx-20xx
		start, err := strconv.Atoi(s[:4])
		if err != nil {
			return nil, err
		}
		end, err := strconv.Atoi(s[5:])
		if err != nil {
			return nil, err
		}
		if start >= end {
			return nil, errors.New("start year must be before end year")
		}
		return getyears(intslice(start, end)), nil
	}
	return nil, errors.New("unrecognized date")
}

func gThan(x, y float64) bool  { return x > y }
func lThan(x, y float64) bool  { return x < y }
func geThan(x, y float64) bool { return x >= y }
func leThan(x, y float64) bool { return x <= y }
func equal(x, y float64) bool  { return x == y }

func (pl playerS) String() string {
	var s string
	for _, p := range pl {
		s = s + p.String() + "\n"
	}

	if len(s) < 2 {
		return "No Players"
	}
	return s[:len(s)-1]
}

func (pl playerS) SmartString(bat, pit bool) string {
	var s string

	for _, p := range pl {
		s = s + fmt.Sprintf("%s\n", p.Bio())
		if bat {
			s = fmt.Sprintf("%sBatting:  %s\n", s, p.Bat)
		}
		if pit {
			s = fmt.Sprintf("%sPitching: %s\n", s, p.Pit)
		}
	}

	if len(s) == 0 {
		return ""
	}
	return s[:len(s)-1]
}
