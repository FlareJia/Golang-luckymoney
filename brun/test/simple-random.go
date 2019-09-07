package main

import (
	"fmt"
	"resk-projects/src/wlarein.com/resk/infra/algo"
)

func Test1main() {
	count, amount := int64(10), int64(100)
	for i:=int64(0); i<count; i++ {
		x := algo.SimpleRand(count, amount*100)
		fmt.Print(float64(x)/float64(100), ",")
	}
	fmt.Println()
}
