package errs

var InternalErr = NewInternalError("произошла ошибка")

type InternalError string

func (e InternalError) Error() string {
	return string(e)
}

func (InternalError) Id() string {
	return "2f611bdf-6a8b-4f3d-bbef-3fe6e0f060cb"
}

func NewInternalError(message string) InternalError {
	return InternalError(message)
}
