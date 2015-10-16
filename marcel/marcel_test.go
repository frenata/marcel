package marcel

import (
	"fmt"
	"testing"

	"github.com/frenata/marcel/lahman"
)

func Test_GetYear(t *testing.T) {
	years := []int{}
	for i := 1871; i < 2015; i++ {
		years = append(years, i)
	}
	lahman.Load(years...)
	count := 0
	for i := 1871; i < 2015; i++ {
		players := lahman.GetPostYear(i)
		for _, p := range players {
			if p.Pit.SHO() > 0 && p.Pit.H() == 0 {
				//fmt.Println(p)
				count++
			}
		}
	}
	if count != 1 {
		t.Log("World series shutouts in history: ", count)
		t.Fatal("there should be only one world series shutout in history!")
	}
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
	p, pa, birth := weightPlayer("beltrca01", 2004)
	t.Log("beltran's batting stats", p.PA(), p)
	t.Log("plate appearances", pa)
	t.Log("birth", birth)

	if p.HR() != 318 {
		t.Fatal("did not weight HR's correctly by year")
	}
	if p.PA() != 7938 {
		t.Fatal("did not weight PA's correctly by year")
	}
}

func Test_regressPlayer(t *testing.T) {
	//regress := regressPlayer("beltrca01", 2004)
	regress := RegressPlayer("beltrad01", 2004)
	fmt.Println(regress.Precise())
}
