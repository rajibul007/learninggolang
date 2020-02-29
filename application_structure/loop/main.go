package main

import "fmt"

func main() {
	c := 0
	i := 1
	for {

		if i%17 == 0 {
			fmt.Println(i)
			c++

		}
		i++
		if c == 20 {
			break
		}
	}

	for a := 1; a <= 20; a++ {
		if a%2 != 0 {

			continue

		}

		fmt.Println(a)

	}

}
