package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for entering your name as , ", input)
	fmt.Printf("Type of this name  is %T", input)
}
