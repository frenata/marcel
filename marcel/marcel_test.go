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
	bat, _ := leagueAvg(2003)
	if fmt.Sprintf("%.4f", bat.HR()) != "0.0285" {
		t.Log(bat.HR())
		t.Fatal("doesn't match TT")
	}
	bat, _ = leagueAvg(2002)
	if fmt.Sprintf("%.4f", bat.HR()) != "0.0279" {
		t.Log(bat.HR())
		t.Fatal("doesn't match TT")
	}
	bat, _ = leagueAvg(2001)
	if fmt.Sprintf("%.4f", bat.HR()) != "0.0300" {
		t.Log(bat.HR())
		t.Fatal("doesn't match TT")
	}
}

func Test_weightPlayer(t *testing.T) {
	//weightPlayer(2004, "beltraca01")
}
