package main_test

import (
	"fmt"
	"testing"

	"github.com/frenata/marcel"
)

func Test_GetYear(t *testing.T) {
	players := main.GetYear(2003)

	count := 0
	for _, p := range players {
		if p.Pitching != nil && p.Pitching.BFP < p.Batting.AB {
			fmt.Println(p)
			count++
		}
	}
	fmt.Println(count)
}
