package main

import (
	"bufio"
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

var playersQ playerS

func init() {
	years := intslice(FirstYear, LastYear)

	lahman.Load(years...)

	playersQ = loadyears(years)
}

func intslice(start, end int) (ns []int) {
	for i := start; i <= end; i++ {
		ns = append(ns, i)
	}
	return ns
}

func loadyears(years []int) (ps playerS) {
	for _, y := range years {
		playerYears := lahman.GetYear(y)
		for _, p := range playerYears {
			ps = append(ps, p)
		}
	}
	return ps
}

func getyears(years []int) (ps playerS) {
	for _, p := range playersQ {
		for _, y := range years {
			if p.Year() == y {
				ps = append(ps, p)
				break
			}
		}
	}
	return ps
}

func main() {
	fmt.Println("BB Query")

	inputLoop()
}

func inputLoop() {
	cliReader := bufio.NewReader(os.Stdin)

exit:
	for {
		fmt.Println("Enter a query.")
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
			//fmt.Println(res)
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
