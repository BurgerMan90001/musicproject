package repository

var _ error = (*Error)(nil)

type Error struct {
	s string
}

func (e *Error) Error() string {
	return e.s
}
func newRepoErr(s string) *Error {
	return &Error{s}
}

var (
	ErrNotFound = newRepoErr("not found")
	ErrNilRepo  = newRepoErr("repo in context is nil")
)
