package main

import "fmt"

type sujon string
type aa = int64

func main() {

	type raj string
	var a raj = "rajibul"
	var b sujon
	b = sujon(a)
	fmt.Println(a, b)
}
