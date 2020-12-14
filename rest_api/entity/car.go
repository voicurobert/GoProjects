package entity

//Car struct
type Car struct {
	CarData `json:"car"`
}

//CarData struct
type CarData struct {
	ID    int    `json:"id"`
	Brand string `json:"car"`
	Model string `json:"car_model"`
	Year  int    `json:"car_model_year"`
}
