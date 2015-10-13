package marcel

import (
	"fmt"
	"testing"

	"github.com/frenata/marcel/lahman"
)

func Test_GetYear(t *testing.T) {
	count := 0
	for i := 0; i < 2015; i++ {
		players := lahman.GetPostYear(i)
		for _, p := range players {
			if p.Pit.SHO() > 0 && p.Pit.H() == 0 {
				//fmt.Println(p)
				count++
			}
		}
	}
	//fmt.Println("World series shutouts in history: ", count)
}

func Test_LeagueAvg(t *testing.T) {
	bat, pit := leagueAvg(2003)
	fmt.Println(bat)
	fmt.Println(pit)

}
