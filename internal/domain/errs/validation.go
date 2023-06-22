package errs

import "encoding/json"

type ValidationErrors struct {
	messages []string
}

func (ValidationErrors) Id() string {
	return "15b957a3-2f4d-497b-b934-27f3936bff5e"
}

func (v *ValidationErrors) addError(err error) {
	v.messages = append(v.messages, err.Error())
}

func (v *ValidationErrors) Error() string {
	e := struct {
		Message []string `json:"error_messages"`
	}{
		Message: v.messages,
	}

	bytes, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}

	return string(bytes)
}

func (v *ValidationErrors) hasErrors() bool {
	return len(v.messages) > 0
}
