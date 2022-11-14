package main

import (
	"fmt"
)


func main() {
	colors := map[string] string {
		"red": "#FF0000",
		"green": "#4BF745",
		"white": "#FFFFFF",
	}

	printMap(colors)

}

func printMap(c map[string]string) {
	// color is the key, hex is the value
	for color, hex := range c {
		fmt.Println("Color: ", color, ", Hex: ", hex)
	}
}
//var emptyColors map[string] string

// newColors := make(map[string]string)

// newColors["white"] = "#FFFFFF" 

// fmt.Println("EmptyColors: " , newColors)

// // map[string] string: all keys are strings, values too
// colors := map[string] string {
// 	"red": "#ff0000",
// 	"green": "#4bf745",
// 	"blue": "#5sf4d3",
// }

// fmt.Println(colors)

// colors := make(map[int]string)

// colors[10] = "#ffffff"

// delete(colors, 10)
// fmt.Println(colors)