package main

import "fmt"

func main() {

	x := 3 //0011
	y := 2 //0010

	var a int = 4
	var b int32 = 7

	var m complex64 = 2 + 3i

	var n complex64 = 4 + 2i

	fmt.Printf("%v\n", x+y)
	fmt.Printf("%v\n", x-y)
	fmt.Printf("%v\n", x*y)
	fmt.Printf("%v\n", x/y)
	fmt.Printf("%v\n", x^y)
	fmt.Printf("%v\n", x%y)

	fmt.Printf("%v\n", a+int(b))

	fmt.Printf("%v\n", x&y)
	fmt.Printf("%v\n", x/y)
	fmt.Printf("%v\n", x|y)
	fmt.Printf("%v\n", x&^y)

	fmt.Printf("%v\n", m+n)
	fmt.Printf("%v\n", m-n)
	fmt.Printf("%v\n", m/n)
	fmt.Printf("%v\n", m*n)

}
