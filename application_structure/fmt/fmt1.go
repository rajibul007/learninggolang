package main

import "fmt"

func main() {
	a := "One"
	b := 2
	c := "ka"
	d := "four"

	fmt.Println("My Name is KHAN and my recent movie is ", a, b, c, d)

	fmt.Printf("My name is Khan my movie name is %s %d %v %v", a, b, c, d)
	fmt.Printf("My name is Khan my movie name is %t %t %t %T", a, b, c, d)

}
