package main

import "fmt"

func main() {
	x := 1

	names := []string{"Swati", "Asha", "Usha", "Rekha"}

	fmt.Println(names)

	for i := 0; i < 5; i++ {
		fmt.Println(i + x)
	}

	for i := 0; i < len(names); i++ {
		//fmt.Println(names)
		fmt.Println(names[i])
	}

	for i := range names {
		fmt.Println(names[i])
	}

	// for x < 10 {
	// 	x += x

	// 	if x == 8 {
	// 		break
	// 	}
	// 	fmt.Println("Now value of x is:", x)
	// }

	//

	for x < 10 {
		x += x

		if x == 8 {
			goto xyz
		}
		fmt.Println("Now value of x is:", x)
	}

xyz:
	fmt.Println("come out of the loop")
}
