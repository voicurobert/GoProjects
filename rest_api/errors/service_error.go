package errors

//ServiceError returns business error messages
type ServiceError struct {
	Message string `json:"message"`
}
