package errors

type CustomError struct {
	status  int64
	message string
}

func New(statusCode int64, messageText string) *CustomError {
	return &CustomError{
		status: statusCode,
		message: messageText,
	}
}
