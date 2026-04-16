package upload

import (
	"testing"

	"musicproject.com/internal/services/encode"
)

func TestUploadMetadata(t *testing.T) {
	//ctx := t.Context()
	encoder := encode.NewFFmpeg()
	New(nil, encoder)
}
