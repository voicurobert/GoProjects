package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/voicurobert/GoProjects/rest_api/entity"
)

//CarDetailsService interface
type CarDetailsService interface {
	GetDetails() entity.CarDetails
}

type carDetailsService struct {
}

var (
	carService       CarService   = NewCarService()
	ownerService     OwnerService = NewOwnerService()
	carDataChannel                = make(chan *http.Response)
	ownerDataChannel              = make(chan *http.Response)
)

//NewCarDetailsService constructor
func NewCarDetailsService() CarDetailsService {
	return &carDetailsService{}
}

func (*carDetailsService) GetDetails() entity.CarDetails {
	// create a go routing to call endpoint 1
	go carService.FetchData()
	// go routine to call endpoint 2
	go ownerService.FetchData()

	car, _ := getCarData()
	owner, _ := getOwnerData()

	return entity.CarDetails{
		ID:        car.ID,
		Brand:     car.Brand,
		Model:     car.Model,
		Year:      car.Year,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,
		Email:     owner.Email,
	}

}

func getCarData() (entity.Car, error) {
	data := <-carDataChannel
	var car entity.Car
	err := json.NewDecoder(data.Body).Decode(&car)
	if err != nil {
		fmt.Println(err.Error())
		return car, err
	}
	return car, nil
}

func getOwnerData() (entity.Owner, error) {
	data := <-ownerDataChannel
	var owner entity.Owner
	err := json.NewDecoder(data.Body).Decode(&owner)
	if err != nil {
		fmt.Println(err.Error())
		return owner, err
	}
	return owner, nil
}
