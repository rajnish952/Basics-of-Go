package main

import (
	"fmt"
	"os"
)

func main() {
	f := os.Args[1]
	fmt.Println(f)
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("Read Error!", err)
		os.Exit(1)
	}
	d := make([]byte, 99999)
	cnt, err := file.Read(d)
	if err != nil {
		fmt.Println("Read Error!! ", err)

	}
	fmt.Printf("%q", d[:cnt])

}
