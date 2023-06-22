package errs

type NotFoundError string

var NotFoundErr = NewInternalError("не найдено")

const NotFoundCode = "ab169ab0-9f2b-4748-aaee-0a14423dffed"

func (e NotFoundError) Error() string {
	return string(e)
}

func (NotFoundError) Id() string {
	return NotFoundCode
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError(message)
}
