package upload

import (
	"errors"

	"musicproject.com/pkg/model"
)

func validateUploadRequest(req *model.UploadSongRequest) error {
	if req == nil {
		return errors.New("Song request is empty")
	}
	var errorList []error
	if req.Filename == "" {
		errorList = append(errorList, errors.New("Filename is empty"))
	}
	if req.Genre == "" {
		errorList = append(errorList, errors.New("Genre is empty"))
	}

	return errors.Join(errorList...)
}
