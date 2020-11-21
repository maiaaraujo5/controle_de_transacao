package errors

type notFound struct {
	Err
}

func NotFound(message string) error {
	return &notFound{
		Err{message: message},
	}
}

func IsNotFound(err error) bool {
	_, ok := err.(*notFound)
	return ok
}
