package main

import (
	"fmt"
	"resk-projects/src/wlarein.com/resk/infra/algo"
)

func main() {
	count, amount :=int64(10), int64(100)

	for i:=int64(0); i<count; i++ {
		x:=algo.DoubleAverage(count, amount*100)
		fmt.Print(x,",")
	}
	println()
}
