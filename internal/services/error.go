package services

var _ error = (*Error)(nil)

type Error struct {
	s string
}

func (e *Error) Error() string {
	return e.Error()
}

func NewErr(s string) *Error {
	return &Error{s}
}

var ()
