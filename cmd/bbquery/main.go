package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/frenata/marcel/lahman"
)

const (
	FirstYear       = 1871
	FirstModernYear = 1900
	LastYear        = 2015
)

type playerS []*lahman.Player

var config struct {
	post  bool
	years []int
	playerS
}

func init() {
	flag.BoolVar(&config.post, "post", false, "set true to query postseason stats")

	years := ""
	flag.StringVar(&years, "years", "", "set years to query")
	flag.Parse()

	var err error
	config.years, err = valYears(years)
	if err != nil && years != "" {
		fmt.Println("'years' flag: ", err)
	}
}

func intslice(start, end int) (ns []int) {
	for i := start; i <= end; i++ {
		ns = append(ns, i)
	}
	return ns
}

func loadyears() (ps playerS) {
	for _, y := range config.years {
		var playerYears playerS
		if config.post {
			playerYears = lahman.GetPostYear(y)
		} else {
			playerYears = lahman.GetYear(y)
		}
		for _, p := range playerYears {
			ps = append(ps, p)
		}
	}
	return ps
}

func getyears() (ps playerS) {
	for _, p := range config.playerS {
		for _, y := range config.years {
			if p.Year() == y {
				ps = append(ps, p)
				break
			}
		}
	}
	return ps
}

func load() {
	if len(config.years) == 0 {
		config.years = intslice(FirstYear, LastYear)
	}
	lahman.Load(config.years...)
	config.playerS = loadyears()
}

func main() {
	load()
	fmt.Println("BB Query")

	inputLoop()
}

func inputLoop() {
	cliReader := bufio.NewReader(os.Stdin)

exit:
	for {
		fmt.Println("\nEnter a query.")
		//fmt.Println(config.years)
		input, err := cliReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		input = strings.TrimSpace(input)

		switch input {
		case "quit":
			fallthrough
		case "q":
			break exit
		default:
			res, err := Query(input)
			if err != nil {
				fmt.Println(err)
				break
			}

			err = pipeLess(res)
			if strings.Contains(fmt.Sprint(err), "file not found") {
				fmt.Println(res)
				fmt.Println("Simple printing: ", err)
			} else if err != nil {
				fmt.Println("Unknown error: ", err)
			}
		}
	}
	fmt.Println("Exiting.")
}
