package main

import "fmt"

type Student struct {
	rollno int
	fname  string
	lname  string
	marks  float32
}

func main() {
	var s1 = Student{
		100,
		"Swati",
		"Naik",
		45.5,
	}
	fmt.Println(s1)

	s2 := Student{110, "usha", "naik", 79}
	fmt.Printf("Roll number of %v is : %v ", s2.fname, s2.rollno)

}

//https://github.com/swatiganesh/all-examples.git
