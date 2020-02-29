package main

import (
	"fmt"
	"os"
)

func main() {

	c := (len(os.Args))

	if c < 2 {
		fmt.Println("please mention any language as a argument")
	} else {

		lang := os.Args[1]
		switch lang {
		case "python":
			fmt.Println("python is awesome ..good for devops")
		case "go", "golang":
			fmt.Println("go is the future")

		default:
			fmt.Println("start learning go or python ")

		}
	}
}
