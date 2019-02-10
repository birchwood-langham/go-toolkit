package errors

// NotImplementedError is an error that can be used as a placeholder to indicated that something
// has yet to be implemented
type NotImplementedError struct{}

func (n NotImplementedError) Error() string {
	return "not implemented"
}

// NotImplemented is a utility function for creating a NotImplementedError
func NotImplemented() error {
	return NotImplementedError{}
}
