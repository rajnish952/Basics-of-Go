package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//fmt.Println(res)
	bs := make([]byte, 99999)

	// Reader Interface
	// Read body of http response and store it in byte slice, as read take byte slice argument to store the content
	res.Body.Read(bs)
	fmt.Println(string(bs))

	// Other way of logging out body - writer interface is implemented by os.Stdout if of type file which implements writer interface
	// -> write []byte
	io.Copy(os.Stdout, res.Body)

}
