package main

import (
	"fmt"

	"github.com/crisdelrosario/timeutil"
)

func main() {
	tu := timeutil.New("Asia/Singapore")
	fmt.Println("Current Date and Time: ", tu.Now())

	tu.Elapse.Start()
	count := 1000000000
	size := count / 100
	x := 0
	y := 0
	for y < count {
		if (y % size) == 0 {
			fmt.Print(".")
		}

		for x < count {
			x++
		}
		y++
	}
	fmt.Println("")

	elapsed := tu.Elapse.Stop()
	measurement := "ms"
	if elapsed > 1000 {
		elapsed = elapsed / 1000
		measurement = "s"
	}

	fmt.Println("Elapsed Time: ", elapsed, measurement)
}
