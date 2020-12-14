package service

import (
	"fmt"
	"net/http"
)

//OwnerService interface
type OwnerService interface {
	FetchData()
}

type fetchOwnerDataService struct {
}

const (
	ownerServiceURL = "https://myfakeapi.com/api/users/1"
)

//NewOwnerService constructor
func NewOwnerService() OwnerService {
	return &fetchOwnerDataService{}
}

func (*fetchOwnerDataService) FetchData() {
	client := http.Client{}
	fmt.Println("Fetching the url &s", ownerServiceURL)

	// call external api
	resp, _ := client.Get(ownerServiceURL)

	//TODO: Write response to the channel
	ownerDataChannel <- resp
}
