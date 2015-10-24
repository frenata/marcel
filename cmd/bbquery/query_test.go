package main

import (
	"fmt"
	"testing"
)

func Test_QueryB(t *testing.T) {
	load()
	testQ := "b HR > 69"
	result, _ := Query(testQ)
	fmt.Println(result)
}
