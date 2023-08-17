package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "#ff0000",
	}

	printMap((colors))

	// // assign zero value
	// var colors2 map[string]string

	// // same as var colors2
	// colors3 := make(map[string]string)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
