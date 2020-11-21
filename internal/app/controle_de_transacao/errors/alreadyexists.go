package errors

type alreadyExists struct {
	Err
}

func AlreadyExists(message string) error {
	return &alreadyExists{
		Err{message: message},
	}
}

func IsAlreadyExists(err error) bool {
	_, ok := err.(*alreadyExists)
	return ok
}
