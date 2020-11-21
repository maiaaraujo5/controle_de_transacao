package errors

type badRequest struct {
	Err
}

func BadRequest(message string) error {
	return &badRequest{
		Err{message: message},
	}
}

func IsBadRequest(err error) bool {
	_, ok := err.(*badRequest)
	return ok
}
