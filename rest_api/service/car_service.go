package service

import (
	"fmt"
	"net/http"
)

//CarService interface
type CarService interface {
	FetchData()
}

type fetchCarDataService struct {
}

const (
	carServiceURL = "https://myfakeapi.com/api/cars/1"
)

//NewCarService constructor
func NewCarService() CarService {
	return &fetchCarDataService{}
}

func (*fetchCarDataService) FetchData() {
	client := http.Client{}
	fmt.Println("Fetching the url &s", carServiceURL)

	// call external api
	resp, _ := client.Get(carServiceURL)

	//Write response to the channel
	carDataChannel <- resp

}
