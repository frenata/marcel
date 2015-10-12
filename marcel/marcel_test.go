package marcel_test

import (
	"fmt"
	"testing"

	"github.com/frenata/marcel/lahman"
)

func Test_GetYear(t *testing.T) {
	count := 0
	for i := 0; i < 2015; i++ {
		players := lahman.GetYear(float64(i))
		for _, p := range players {
			if p.Pit.SHO() > 0 && p.Pit.H() == 0 {
				fmt.Println(p)
				count++
			}
		}
	}
	fmt.Println("World series shutouts in history: ", count)
}
