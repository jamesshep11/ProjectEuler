package main

import "fmt"

func main() {
	for c := 1000; c > 3; c-- {
		for b := c - 1; b > 2; b-- {
			for a := b - 1; a > 1; a-- {
				if a+b+c != 1000 {
					continue
				}

				if (a*a)+(b*b) != (c * c) {
					continue
				}

				fmt.Println(a, b, c, a*b*c)
				return
			}
		}
	}

	fmt.Println("Not found")
}
