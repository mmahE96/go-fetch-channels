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
		"http://yahoo.com",
	}
	c := make(chan string)

	for _, link := range links {

		go makeRequest(link, c)

	}
	for {
		go makeRequest(<-c, c)
	}

}

func makeRequest(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link

}
