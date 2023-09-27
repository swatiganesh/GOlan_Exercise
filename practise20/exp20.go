package main

import "fmt"

func main() {
	defer fmt.Println("World") // 5th
	defer fmt.Println("One")   // 4th
	defer fmt.Println("Two")   // 3rd
	fmt.Println("Hello")       // 1st execute  after this main will going to call func myDefer
	myDefer()                  //2nd

}

// world, One, Two
// 0, 1, 2, 3, 4
// hello, 43210, two, One, world

func myDefer() {
	for i := 0; i < 5; i++ { // 4
		// 3
		// 2
		// 1
		defer fmt.Print(i, "\n") // here we ate using defer so act like a 1st in last out so it will print in reverse order
		//like 4, 3, 2, 1 and then again go back to the main and check for further task there also we have defer statements
		// so it will print from the last statement like, two, one, world
	}
}
