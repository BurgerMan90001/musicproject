package handler

// var _ http.ResponseWriter = (*Writer)(nil)

// type Writer struct {
// 	http http.ResponseWriter
// }

// func (w *Writer) Header() http.Header {
// 	return w.http.Header()
// }

// func (w *Writer) Write(b []byte) (int, error) {
// 	return w.http.Write(b)
// }

// // Custom http messages
// func (w *Writer) WriteHeader(code int, message string) {
// 	jerr := &model.Error{}
// 	switch code {
// 	case http.StatusNotFound:
// 		jerr.Code = code
// 		if message == "" {
// 			jerr.Message = "Resource not found"
// 		}
// 		jsonutil.WriteJSON(w.http, jerr, code)
// 	case http.StatusMethodNotAllowed:

// 	}
// }
