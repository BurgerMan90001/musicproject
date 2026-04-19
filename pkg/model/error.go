package model

var _ error = (*Error)(nil)

type Error struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}

func (e *Error) Error() string {
	return e.Message
}

