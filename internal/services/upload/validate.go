package upload

import (
	"net/http"

	"songsled.com/pkg/model"
)

func validateUploadRequest(req *model.SongUploadRequest) error {
	// var err = &model.Error{
	// 	Code:    http.StatusBadRequest,
	// 	Message: "Invalid upload song request body",
	// }
	if req == nil {
		return &model.Error{
			Code:    http.StatusBadRequest,
			Message: "Song request is empty",
		}
	}

	if req.Name == "" {
		return &model.Error{
			Code:    http.StatusBadRequest,
			Message: "Name is empty",
		}
	}
	if len(req.Artists) == 0 {
		return &model.Error{
			Code:    http.StatusBadRequest,
			Message: "No artists in song request",
		}
	}
	// Any errors apear
	// if len(err.Details) > 0 {
	// 	return err
	// }

	return nil
}
