package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
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

	wg.Done()
}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}

	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
		fmt.Println(strings.Repeat("#", 20))
	}

	fmt.Println("No of go routines", runtime.NumGoroutine())

	wg.Wait()
}
