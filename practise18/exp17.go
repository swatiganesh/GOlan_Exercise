package practise18

import (
	"fmt"
	"regexp"
)

func Regex() {

	// string in which the pattern
	// is to be found
	str := "I like jamoon"

	// returns true if the pattern is present
	// in the string, else returns false
	// err is nil if the regexp is valid
	match1, err := regexp.MatchString("jam", str)
	fmt.Println("Match: ", match1, " Error: ", err)

	// this returns false as the match
	// is unsuccessful
	str2 := "ComputerScience"
	match2, err := regexp.MatchString("like", str2)
	fmt.Println("Match: ", match2, "Error: ", err)

	// this will throw an error
	// as the pattern is not valid
	match3, err := regexp.MatchString("lik(e", str2)
	fmt.Println("Match: ", match3, "Error: ", err)
}
