package main

import "fmt"

var f = false

func main() {
	var a = "apple"
	var b, c = 1, 2
	d := 3
	d, e := 5, 6
	var z string
	z = "2"
	//var e = true
	b, c = c, b
	/*var (
		x int
		y int32
		z float64
	)
	*/

	//_, _, _ = x, y, z
	var sum = (2 + 3)
	sub := (100 - 89)
	la, _ := strconv.ParseIntint64(z, 0, 64)
	fmt.Println(a, b, c, d, e, f, la)
	fmt.Println("sum is", sum, "subtracte is ", sub)

}
