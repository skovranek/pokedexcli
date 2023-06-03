package main

import (
	"fmt"
)

func getIndex(order int) string {
	index := fmt.Sprint(order)
	zeros := ""

	if order < 10 {
		zeros = "00"
	} else if order < 100 {
		zeros = "0"
	}
	return (zeros + index)
}