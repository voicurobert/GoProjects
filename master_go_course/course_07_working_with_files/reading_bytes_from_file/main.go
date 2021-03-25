package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := make([]byte, 2)
	nrBytesRead, err := io.ReadFull(file, byteSlice)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Nr of bytes read: %d \n", nrBytesRead)
	log.Printf("%s \n", byteSlice)

	file, err = os.Open("linked_list.go")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)

	data, err = ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}
