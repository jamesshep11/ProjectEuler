package main

import "fmt"

var primes []int

func main() {
	var sum int = 2

	for i := 3; i < 2000000; i += 2 {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}

func isPrime(num int) bool {
	for i := 2; i < 10 && i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	for _, prime := range primes {
		if num%prime == 0 {
			return false
		}
	}

	if num > 9 {
		primes = append(primes, num)
	}

	return true
}
