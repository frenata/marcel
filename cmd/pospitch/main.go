package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/frenata/marcel/lahman"
	"github.com/frenata/marcel/marcel"
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
	pp := []*lahman.Pitcher{}

	for i := min; i <= max; i++ {
		players := marcel.GetYear(int16(i))
		for _, p := range players {
			if p.Pitching != nil && p.Pitching.BFP < p.Batting.AB {
				pp = append(pp, p.Pitching)
				count++
				apps += p.Pitching.G
			}
		}
	}

	for _, p := range pp {
		fmt.Println(p)
	}
	fmt.Printf("Position players pitching: %d\n", count)
	fmt.Printf("Position players pitched %d times total.\n", apps)
}
