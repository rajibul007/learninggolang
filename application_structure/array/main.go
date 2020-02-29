package main

import "fmt"

func main() {

	//declaring array

	var a [3]int              //type 1
	b := [...]int{1, 2, 3, 4} //type 2
	c := [2]int{1, 1}         //type 3
	d := [3]string{"raj", "wasim", "aamir"}
	fmt.Printf("%#v\n", d)
	fmt.Printf("%#v\n%#v\n%#v\n", a, b, c)
	fmt.Println(len(b)) //lenth func
	a[2] = 5            //modifying array content
	fmt.Println(a, d)

	// itteration using array 1st way using range
	e := [...]int{10, 11, 12, 13, 14, 15}
	for i, j := range e { //// range is a language keyword used for iteration
		fmt.Println("index:", i, "number:", j)
	}
	// iterratin java c style

	for i := 0; i < len(e); i++ {

		fmt.Println("index:", i, "number:", e[i])
	}
	//multi dimensional array

	aaa := [2][3]int{
		{1, 2, 3},
		{2, 3, 4},
	}
	for _, i := range aaa {
		for _, j := range i {
			fmt.Printf("%d", j)
		}
	}

	n := [...]int{22, 22, 22}
	m := n
	fmt.Println()
	m[2] = 23
	fmt.Println(m == n)

}
