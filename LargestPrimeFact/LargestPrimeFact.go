package main

import "fmt"

func main() {
	val := 600851475143
	half := val / 2
	var max int

	for i := 1; i <= half; i++ {
		if val%i == 0 {
			j := val / i

			if i > j {
				fmt.Print(max)
				return
			}

			if isPrime(j) {
				fmt.Print(j)
				return
			}

			if isPrime(i) {
				max = i
			}
		}
	}

	fmt.Print(max)
}

func isPrime(num int) bool {
	half := num / 2

	for i := 2; i < half; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
