package errs

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

var (
	CodeNotFound = "not_found"
)

type ErrorSt struct {
	Code     string
	Resource string
	Field    string
}

func (*ErrorSt) Error() string {
	return "tekitou"
}
