package model

var _ error = (*Error)(nil)

type Error struct {
	// Required
	// Status code of the error
	Code int `json:"code"`
	// Required
	// Main error message to show
	Message string `json:"message"`
	// Optional
	// More detailed errors for debugging
	Details string `json:"details,omitempty"`
}

func (e *Error) Error() string {
	return e.Message
}
