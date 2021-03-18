package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

func checkAndSaveBody(url string, c chan string) {
	response, err := http.Get(url)

	if err != nil {
		s := fmt.Sprintf("%s is down\n", url)
		s += fmt.Sprintf("Error: %v \n", err)
		c <- s
	} else {
		defer response.Body.Close()
		s := fmt.Sprintf("%s status code %d \n", url, response.StatusCode)

		if response.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
			}
			file := strings.Split(url, "//")[1] // http://www.google.com
			file += ".txt"
			s += fmt.Sprintf("Writing respond body to %s \n", file)
			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				s += "Error writing file"
				c <- s
			}
			s += fmt.Sprintf("%s is up \n", url)
		}
	}
}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}

	c := make(chan string)

	for _, url := range urls {
		go checkAndSaveBody(url, c)
		fmt.Println(strings.Repeat("#", 20))
	}
	fmt.Println("No of goroutines", runtime.NumGoroutine())

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}
