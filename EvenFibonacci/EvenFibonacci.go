package main

import "fmt"

func main() {
	doEvenFib(1, 1, 0)
}

func doEvenFib(a int, b int, sum int) {
	new := a + b

	if new%2 == 0 {
		sum += new
	}

	if new > 4000000 {
		fmt.Print(sum)
	} else {
		doEvenFib(b, new, sum)
	}
}
