package errors

type CustomError struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`
}

func New(statusCode int64, messageText string)* CustomError {
	return &CustomError{
		Status: statusCode,
		Message: messageText,
	}
}
