package errs

type InvalidArgumentError string

const InvalidArgumentCode = "15b957a3-2f4d-497b-b934-27f3936bff5e"

func (e InvalidArgumentError) Error() string {
	return string(e)
}

func (InvalidArgumentError) Id() string {
	return InvalidArgumentCode
}

func NewInvalidArgumentError(message string) InvalidArgumentError {
	return InvalidArgumentError(message)
}
