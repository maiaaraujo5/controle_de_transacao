package errors

type Err struct {
	message string
}

func (e *Err) Error() string {
	return e.message
}
