package main

import "fmt"

var num int = 65536

func main() {
	var n uint16

	n = uint16(num)

	fmt.Println(n)

}
