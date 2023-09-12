package main

import "fmt"

func main() {
	var i int = 20

	for true {
		found := true

		for j := 2; j <= 20; j++ {
			if i%j != 0 {
				found = false
				break
			}
		}

		if found {
			fmt.Println(i)
			return
		}

		i += 10
	}
}
