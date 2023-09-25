package main

import "fmt"

// func add(x int, y int) int {     //here you can write (x, y int),in  this case x will take the data type which is recently assigned to next value i.e int

// 	var sum = x + y
// 	return sum
// }

// func main() {
// 	num1 := 4

// 	num2 := 6

// 	result := add(num1, num2)

// 	fmt.Println(result)
// }

//in go we can return 2 or more values

func calc(x, y int) (int, int) { // here ur returning 2 values so need to specify the data type of the same
	//func calc(x, y int)(sum,sub, int) {       alternative method to return the value
	//  sum = x +y
	// sub = x-y
	//return
	//}

	var sum = x + y
	var sub = x - y
	return sum, sub
}

func main() {
	a := 2

	b := 3

	result1, result2 := calc(a, b)

	fmt.Println(result1, "\n", result2)
}
