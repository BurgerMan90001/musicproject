package upload

import (
	"net/http"

	"songsled.com/pkg/model"
)

func validateUploadRequest(req *model.UploadSongRequest) error {
	var err = &model.Error{
		Code:    http.StatusBadRequest,
		Message: "Invalid upload song request body",
	}
	if req == nil {
		return &model.Error{
			Code:    http.StatusBadRequest,
			Message: "Song request is empty",
		}
	}
	if req.Name == "" {
		err.Details = "Name is empty"
	}
	if req.Filename == "" {
		err.Details = "Filename is empty"
	}
	if req.Genre == "" {
		err.Details = "Genre is empty"
	}

	// Any errors apear
	if len(err.Details) > 0 {
		return err
	}

	return nil
}
