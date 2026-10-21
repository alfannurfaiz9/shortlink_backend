package dto

type Response struct {
	Success bool
	Message string
	Result  any
}

type ErrorResponse struct {
	Success bool
	Message string
}
