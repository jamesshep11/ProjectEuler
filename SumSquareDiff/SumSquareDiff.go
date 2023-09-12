package main

import (
	"fmt"
	"math"
)

func main() {
	num := 100

	result := (float64)(sumSqrs(num) - sqrSum(num))
	temp := math.Abs(result)

	fmt.Printf("%.0f", temp)
}

func sumSqrs(num int) int {
	var sum int

	for i := 1; i <= num; i++ {
		sum += i * i
	}

	return sum
}

func sqrSum(num int) int {
	var sum int

	for i := 1; i <= num; i++ {
		sum += i
	}

	return sum * sum
}
