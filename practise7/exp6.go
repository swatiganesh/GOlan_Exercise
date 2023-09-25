package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("good morning")

	case t.Hour() > 12 && t.Hour() < 7:
		fmt.Println("good evening")

	case t.Hour() > 7 && t.Hour() < 24:
		fmt.Println("good night")
	}
}
