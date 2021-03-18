package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkUrl(url string, c chan string) {
	response, err := http.Get(url)

	if err != nil {
		s := fmt.Sprintf("%s is down\n", url)
		s += fmt.Sprintf("Error: %v \n", err)
		fmt.Println(s)
		c <- url
	} else {
		s := fmt.Sprintf("%s status code %d \n", url, response.StatusCode)
		s += fmt.Sprintf("%s is up \n", url)
		c <- url
	}
}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}

	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
	}

	//for {
	/*
		for url := range c {
			time.Sleep(time.Second * 2)
			go checkUrl(url, c)
			fmt.Println(strings.Repeat("#", 30))
		}
	*/

	for url := range c {
		go func(u string) {
			time.Sleep(time.Second * 2)
			checkUrl(u, c)
		}(url)
	}
}
