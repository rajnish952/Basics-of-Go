package main

import "fmt"

func main() {
	// map [key type]val type
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	/* Other ways to create empty map */
	// var color map[string]string
	// colors := make(map[string]string)
	// colors["red"] = "#ff0000"

	// delete value from map
	//delete(colors, "red")

	//fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for col, hex := range c {
		fmt.Println("Color value :", col, " Hex value:", hex)
	}
}
