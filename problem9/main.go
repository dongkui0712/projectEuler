package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	var a, b, c int
	var limit int = 1000

	var times int
	for a = 1; a < limit/3; a++ {
		for c = int(limit / 3); c < limit/2; c++ {
			b = limit - a - c

			if a*a+b*b == c*c {
				fmt.Printf("[%d]^2 + [%d]^2 = [%d]^2\n", a, b, c)
			}

			times++
		}
	}
	end := time.Now().UnixNano()
	fmt.Println("time = ", (end-start)/1000)
	fmt.Println("times: ", times)
}
