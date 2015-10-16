package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/frenata/marcel/lahman"
)

func main() {
	min, max := 2003, 2003
	if len(os.Args) > 2 {
		min, _ = strconv.Atoi(os.Args[1])
		max, _ = strconv.Atoi(os.Args[2])
	} else if len(os.Args) > 1 {
		min, _ = strconv.Atoi(os.Args[1])
		max = min
	}
	var years []int
	for i := min; i <= max; i++ {
		years = append(years, i)
	}

	players, count, apps := check(min, max)
	for _, p := range players {
		fmt.Println(p)
	}
	fmt.Printf("Position players pitching: %d\n", count)
	fmt.Printf("Position players pitched %d times total.\n", apps)
}

func check(min, max int) (pp []*lahman.Player, count int, apps int) {
	//pp := []*lahman.Player{}

	for i := min; i <= max; i++ {
		players := lahman.GetYear(i)
		for _, p := range players {
			//if p.Pitching != nil && p.Pitching.BFP < p.Batting.AB {
			if p.IsPosPitcher() {
				pp = append(pp, p)
				count++
				apps += int(p.Pit.G())
			}
		}
	}
	return pp, count, apps
}
