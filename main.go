package main

import (
	"fmt"

	"github.com/frenata/marcel/marcel"
)

func main() {
	regress := marcel.RegressPlayer("beltrad01", 2004)
	fmt.Println(regress.Precise())
}
