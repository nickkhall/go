package errors

type CustomError struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`
}

