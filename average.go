package main

import (
	"fmt"

	"github.com/frenata/marcel/marcel"
)

func League(year float64) (float64, float64) { // (bs lahman.BatStats) { //, ps lahman.PitchStats) {
	players := marcel.GetYear(year)
	var HR, PA, countB, countP float64

	for _, p := range players {
		if p.Pitching == nil || marcel.IsPosPitcher(p) {
			countB++
			pa := p.Batting.AB + p.Batting.BB + p.Batting.HBP + p.Batting.SF + p.Batting.SH
			PA += float64(pa)
			HR += float64(p.Batting.HR)
		}
	}

	//fmt.Println(HR, countB)
	HR = HR / countB
	PA = PA / countB
	_ = countP

	return HR, PA
}

func Player(id string, year float64) (float64, float64) {
	players := marcel.GetYear(year)
	var HR, PA float64

	for _, p := range players {
		if p.Batting.ID == id && p.Batting.Year == year {
			HR += p.Batting.HR
			pa := p.Batting.AB + p.Batting.BB + p.Batting.HBP + p.Batting.SF + p.Batting.SH
			PA += pa
		}
	}

	return HR, PA
}

func main() {
	test := "beltrca01"
	_ = test

	hr1, pa1 := League(2003)
	hr2, pa2 := League(2002)
	hr3, pa3 := League(2001)

	fmt.Println(hr1 / pa1)
	fmt.Println(hr2 / pa2)
	fmt.Println(hr3 / pa3)

	bhr1, bpa1 := Player(test, 2003)
	bhr2, bpa2 := Player(test, 2002)
	bhr3, bpa3 := Player(test, 2001)

	fmt.Println(bhr1, bpa1)
	fmt.Println(bhr2, bpa2)
	fmt.Println(bhr3, bpa3)
}
