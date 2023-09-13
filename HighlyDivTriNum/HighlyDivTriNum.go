package main

import (
	"fmt"
	"math"
)

func main() {
	target := 500
	sum := 0

	for i := 1; true; i++ {
		sum += i
		factors := getFactors(sum)

		if len(factors) > target {
			fmt.Println(sum)
			return
		}
	}
}

func getFactors(num int) []int {
	var factors []int
	sqrt := int(math.Sqrt(float64(num)))

	for i := 1; i <= sqrt; i++ {
		if num%i == 0 {
			factors = append(factors, i)
			factors = append(factors, num/i)
		}
	}

	return factors
}
