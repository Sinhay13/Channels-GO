package main

import (
	"fmt"
	"net/http"
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
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
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
		c <- link + " might be down!"
		return
	}

	// If the request is successful, the link is up
	// Print a message to the console
	fmt.Println(link, "is up!")
	// Send a message to the channel
	c <- link + " is up!"
}
