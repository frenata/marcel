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

	count := 0
	var apps int16 = 0
	pp := []*lahman.Player{}

	for i := min; i <= max; i++ {
		players := lahman.GetYear(i)
		for _, p := range players {
			//if p.Pitching != nil && p.Pitching.BFP < p.Batting.AB {
			if p.IsPosPitcher() {
				pp = append(pp, p)
				count++
				apps += int16(p.Pit.G())
			}
		}
	}

	for _, p := range pp {
		fmt.Println(p)
	}
	fmt.Printf("Position players pitching: %d\n", count)
	fmt.Printf("Position players pitched %d times total.\n", apps)
}
