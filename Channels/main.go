package main

import (
	"fmt"
	"net/http"
)

func main() {

	urls := []string{"http://youtube.com/", "http://www.google.com", "http://apple.com/", "http://linkedin.com/", "http://en.wikipedia.org/"}

	c := make(chan string) // Create channel of type string

	for _, u := range urls {
		//fmt.Println("URL", i+1, ": ", u)
		go checkStat(u, c) // Run go routine

	}
	//fmt.Println(<-c) // Blocking code, waiting for channel output
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c) // Blocks for channel outcome and then proceeds further
	}
}

func checkStat(link string, c chan string) { // c is channel of type string
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " Webseite be down!")
		c <- "Might be Down!" // Send data to channel
		return
	}

	fmt.Println(link, " is UP!")
	c <- "Link is UP!"

}
