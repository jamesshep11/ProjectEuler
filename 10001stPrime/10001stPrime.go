package main

import "fmt"

var primes []int = []int{2}

func main() {
	target := 10001
	num := 3

	for len(primes) < target {
		if isPrime(num) {
			primes = append(primes, num)
		}

		num += 2
	}

	fmt.Print(primes[target-1])
}

func isPrime(num int) bool {
	half := num / 2

	for i := 2; i < 10 && i <= half; i++ {
		if num%i == 0 {
			return false
		}
	}

	for _, prime := range primes {
		if num%prime == 0 {
			return false
		}
	}

	return true
}
