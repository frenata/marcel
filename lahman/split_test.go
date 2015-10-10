package lahman_test

import (
	"testing"

	"github.com/frenata/marcel/lahman"
)

func Test_BattingYear(t *testing.T) {
	batters, _ := lahman.ReadAll("data/Batting.csv")
	res := lahman.BattingYear(2014, batters)

	//for _, r := range res {
	//		fmt.Println(r)
	//}
	if len(res) != 1435 {
		t.Fatal("length of 2014 batters is not 1435")
	}

	//	if res[len(res)-1] != tail {
	//		t.Fatal("Last record does not match tail")
	//	}
}
