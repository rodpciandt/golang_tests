package main

import (
	"fmt"
)



func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {

		if number % 2 == 0 {
			fmt.Printf("%v is Even", number)
		} else {
			fmt.Printf("%v is Odd", number)
		}

		fmt.Println("")
	}
}