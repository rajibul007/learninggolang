package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	age, _ := strconv.ParseInt(os.Args[1], 0, 64)

	fmt.Println(os.Args[2])

	if age > 0 && age < 18 {
		fmt.Printf("you can not vote ...come after %T years", 18-age)
	} else if age == 18 {

		fmt.Printf("This is your first vote ...Enjoy")

	} else if age > 18 && age < 100 {
		fmt.Println("you can vote ")
	} else {
		fmt.Println("invalid")
	}

}
