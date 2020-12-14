package controller

import (
	"encoding/json"
	"net/http"

	"github.com/voicurobert/GoProjects/rest_api/service"
)

//CarDetailsController interface
type CarDetailsController interface {
	GetCarDetails(rw http.ResponseWriter, r *http.Request)
}

type controller struct {
}

var carDetailsService service.CarDetailsService

//NewCarDetailsController constructor
func NewCarDetailsController(cds service.CarDetailsService) CarDetailsController {
	carDetailsService = cds
	return &controller{}
}

func (*controller) GetCarDetails(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("content-type", "application/json")
	result := carDetailsService.GetDetails()
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(result)
}
