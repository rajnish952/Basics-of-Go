package main

import (
	"fmt"
	"math"
)

func basicSyntax() {

	fmt.Println("Go" + "Lang")
	fmt.Println(true && false)
}

func variableDec() {
	var (
		x   int
		str = "Hey there!"
		y   = 10
	)

	c := "Go"

	fmt.Println("X: ", x)
	fmt.Println("Y: ", y)
	fmt.Println("Str: ", str)
	fmt.Println("C: ", c)
}

func costants() {
	const a = 10 // declare const using keyword, numeric const has no type untill given one
	fmt.Println(math.Sin(a))
	fmt.Println("A: ", a)
}
