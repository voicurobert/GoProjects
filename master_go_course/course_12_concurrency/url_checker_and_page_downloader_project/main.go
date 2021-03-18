package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func checkAndSaveBody(url string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is down\n", url)
	} else {
		defer response.Body.Close()
		fmt.Printf("%s status code %d \n", url, response.StatusCode)

		if response.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
			}
			file := strings.Split(url, "//")[1] // http://www.google.com
			file += ".txt"
			fmt.Printf("Writing respond body to %s \n", file)
			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}
	for _, url := range urls {
		checkAndSaveBody(url)
		fmt.Println(strings.Repeat("#", 20))
	}
}
