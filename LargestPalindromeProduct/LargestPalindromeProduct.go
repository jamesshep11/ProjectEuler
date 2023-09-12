package main

import (
	"fmt"
	"strconv"
)

func main() {
	var max int

	for i := 999; i > 99; i-- {
		for j := i; j > 9; j-- {
			num := i * j

			if isPalindrome(num) && num > max {
				max = num
				break
			}
		}
	}

	fmt.Print(max)
}

func isPalindrome(num int) bool {
	sNum := strconv.Itoa(num)
	length := len(sNum) - 1

	for i := 0; i < length-i; i++ {
		if sNum[i] != sNum[length-i] {
			return false
		}
	}

	return true
}
