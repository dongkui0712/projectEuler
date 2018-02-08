/*
Multiples of 3 and 5

If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.
*/
package main

import (
	"fmt"
)

func main() {
	min, max := 0, 1000
	args := []int{3, 5}

	results := SumOfNumThatIsMultiplesOf(min, max, args...)
	fmt.Println("results: ", results)
}

func SumOfNumThatIsMultiplesOf(min, max int, args ...int) (sum int) {

	for i := min; i < max; i++ {
		ok := false
		for _, arg := range args {
			if i%arg == 0 {
				ok = true
				break
			}
		}
		if ok {
			sum += i
		}
	}
	return
}
