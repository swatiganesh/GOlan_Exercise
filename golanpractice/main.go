package main

import "fmt"

func main() {
	//fmt.Println("HI Everyone")

	var fruit = "banana"
	const quantity = 10

	fmt.Println("I like", fruit, "so I bought", quantity, fruit)

	var fname string = "Swati"
	var lname string = "Naik"

	fmt.Printf("\nmy sweet name is %v %v and everyone likes me ", fname, lname)

	fmt.Printf("data type of the fname is %T and quantity is %T ", fname, quantity)
}
