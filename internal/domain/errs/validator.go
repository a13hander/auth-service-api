package errs

import "context"

type Condition func(ctx context.Context) error

func Validate(ctx context.Context, conds ...Condition) error {
	ve := &ValidationErrors{}

	for _, c := range conds {
		err := c(ctx)
		if err != nil {
			if de, ok := AsDomainError(err); ok {
				ve.addError(de)
				continue
			}

			return err
		}
	}

	if ve.hasErrors() {
		return ve
	}

	return nil
}
