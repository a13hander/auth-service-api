package errs

type DomainError interface {
	error
	Id() string
}

func AsDomainError(err error) (DomainError, bool) {
	de, ok := err.(DomainError)

	return de, ok
}
