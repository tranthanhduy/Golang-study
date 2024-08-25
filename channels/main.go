package main

import (
	"fmt"
	"net/http"
	"time"
)

var links []string

func main() {
	links = []string{
		"http://google.com",
		"http://amazon.com",
		"http://golang.org",
		"http://instagram.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// whenever l got value from channel c, it will pass the value to checkLink()
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	resp, err := http.Get(link)

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Link maybe down!")
		c <- link
		return
	}

	fmt.Println("Status of", link, "is", resp.Status)
	fmt.Println("Yes, it's up!")
	// pass the link to channel
	c <- link

}
