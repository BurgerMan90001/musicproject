package handler

var _ error = (*HandlerErr)(nil)
var (
// ErrInternalServerError = &model.Error{
// 	Code:    http.StatusInternalServerError,
// 	Message: "Internal server error",
// }
// ErrInvalidRequestBody = &model.Error{
// 	Code:    http.StatusBadRequest,
// 	Message: "Invalid request body",
// }
)

type HandlerErr struct {
	s string
}

func (e *HandlerErr) Error() string {
	return e.s
}
