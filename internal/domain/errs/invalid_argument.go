package errs

type InvalidArgumentError string

func (e InvalidArgumentError) Error() string {
	return string(e)
}

func (InvalidArgumentError) Id() string {
	return "15b957a3-2f4d-497b-b934-27f3936bff5e"
}

func NewInvalidArgumentError(message string) InvalidArgumentError {
	return InvalidArgumentError(message)
}
