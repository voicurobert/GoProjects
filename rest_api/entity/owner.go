package entity

//Owner struct
type Owner struct {
	OwnerData `json:"User"`
}

//OwnerData struct
type OwnerData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
