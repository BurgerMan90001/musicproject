package integration

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"songsled.com/internal/jsonutil"
	"songsled.com/pkg/model"
)

func (s *testSuite) TestUploadSongMetadata() {
	endpoint := "/v1/songs/upload"

	// token, err := s.jwtAccess.GenerateToken(uuid.Nil)
	// s.Require().NoError(err)

	type test struct {
		name      string
		audioFile string
		imageFile string
		wantCode  int
		req       *request
	}

	tests := []test{
		{
			name:      "invalid file type",
			audioFile: filepath.Join("testdata", "cool-pic-128.jpg"),
			wantCode:  http.StatusBadRequest,
			req: &request{
				method: "PUT",
				body:   map[string]any{},
			},
		},
		{
			name:     "no file",
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			w := s.newRequest(endpoint, tt.req)
			s.Equal(tt.wantCode, w.Code)
			s.Empty(w.Result().Header.Get("Location"))
		})
	}

	s.Run("success", func() {
		tt := test{
			req: &request{
				method: http.MethodPut,
				body: map[string]any{
					"name":         "Cool music",
					"artists":      []string{"rockguy", "rockguy2"},
					"genres":       []string{"Rock", "Pop"},
					"filename":     "mysong.mp3",
					"duration":     123,
					"creationDate": "2006-06-25",
				},
			},
			audioFile: filepath.Join("testdata", "Iwan Gabovitch - Dark Ambience Loop.ogg"),
			imageFile: filepath.Join("testdata", "cool-pic-128.jpg"),

			wantCode: http.StatusOK,
		}

		w := s.newRequest("/v1/audio", tt.req)
		s.Equal(http.StatusContinue, w.Code)

		b, err := jsonutil.ReadJson[model.FileUploadResponse](w.Body)
		s.Require().NoError(err)
		
		// The file location uses the public url
		s.Contains(b.Href, s.cfg.File.Public)

		if !testing.Short() && os.Getenv("USE_CLOUD") == "true" {
			// err := s.uploadFile(tt.audioFile, "")
			// s.Require().NoError(err)

			if tt.imageFile != "" {

				// w2 := s.newRequest(fmt.Sprintf("/v1/songs/%w/upload/cover", songId), &request{
				// 	method: "PUT",
				// })

				// location := w2.Result().Header.Get("Location")
				// err := s.uploadFile(imagePath, location)
				// s.Require().NoError(err)

			}
		}

		w2 := s.newRequest(endpoint, tt.req)

		// location := w2.Result().Header.Get("Location")

		s.Equal(tt.wantCode, w2.Code)
		// s.NotEmpty(location)

		r, err := jsonutil.ReadJson[model.FileUploadResponse](w2.Body)
		s.Require().NoError(err)

		if w2.Result().StatusCode >= 500 && w2.Code <= 600 {

			// Show the error in testing
			s.Equal(r, "asdasd")
		}

		// Check the uploaded song metadata
		// metadataEndpoint, err := url.JoinPath("/v1/songs")
		// s.Require().NoError(err)

		// w3 := s.newRequest(metadataEndpoint, &request{
		// 	method: "GET",
		// })

		// s.Equal(http.StatusOK, w3.Code)

	})
}

func (s *testSuite) uploadFile(path, endpoint string) error {
	file, err := os.Open(path)
	s.Require().NoError(err)

	defer file.Close()

	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)
	part, err := w.CreateFormFile("fileUpload", filepath.Base(path))
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	r, err := http.NewRequest("PUT", endpoint, body)
	if err != nil {
		return err
	}

	r.Header.Set("Content-Length", strconv.Itoa(body.Cap()))

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	s.Equal("", string(data), strconv.Itoa(body.Cap()))

	return nil
}
