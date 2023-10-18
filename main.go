package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// create a channel to receive the results
	c := make(chan string)

	// start a goroutine for each link
	for _, link := range links {
		go checkLink(link, c)
	}

	// receive the results from the channel
	//first way with seconde message
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)

	// }

	// second way
	// infinite loop, but do not forget the second argument "c", "l" is the same like for {go checkLink(<-c,c)}
	// for l := range c {
	// 	go checkLink(l, c)

	//}

	// final way with time sleeping lambda (anonymous) style function
	// The time.Sleep() function is implemented inside the anonymous function rather than directly in the loop
	// to prevent the loop from blocking and to allow other goroutines to execute while the sleep is in progress.
	// This ensures that the program remains responsive and doesn't hang while waiting for the sleep to complete.
	// By using an anonymous function, we can create a separate goroutine for each link that needs to be checked,
	// which allows us to check multiple links concurrently and improves the overall performance of the program.
	for l := range c {
		go func(link string) { // func literal = anonymous func or lambda func
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

// checkLink checks if a given link is up or down by making an HTTP GET request to it.
// If the request returns an error, it means the link is down, and the function sends a message to the channel.
// If the request is successful, the link is up, and the function sends a message to the channel.
func checkLink(link string, c chan string) {
	// Make an HTTP GET request to the link
	_, err := http.Get(link)

	// If the request returns an error, the link is down
	if err != nil {
		// Print a message to the console
		fmt.Println(link, "might be down!")
		// Send a message to the channel
		c <- link //+ " might be down I think!" second message
		return
	}

	// If the request is successful, the link is up
	// Print a message to the console
	fmt.Println(link, "is up!")
	// Send a message to the channel
	c <- link //+ " Yep up!" second message
}
