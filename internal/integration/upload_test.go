package integration

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"songsled.com/internal/jsonutil"
	"songsled.com/pkg/model"
)

func (s *testSuite) TestUploadCover() {
	endpoint := "/v1/images/covers"
	type test struct {
		name      string
		wantCode  int
		imageFile string
		req       *request
	}

	tests := []test{
		{
			imageFile: filepath.Join("testdata", "cool-pic-128.jpg"),
			req: &request{
				method: http.MethodPut,
				body: map[string]any{
					"filename": "cool-pic-128.jpg",
				},
			},
			wantCode: http.StatusContinue,
		},
	}
	for _, tt := range tests {
		w := s.newRequest(endpoint, tt.req)

		b, err := jsonutil.ReadJson[model.FileUploadResponse](w.Body)
		s.Assert().NoError(err)

		s.Equal(tt.wantCode, w.Code)
		s.NotEmpty(b.Links[0].Href)
		s.Equal("upload", b.Links[0].Rel)

	}

}

func (s *testSuite) TestUploadAudio() {
	endpoint := "/v1/audio/songs"

	type test struct {
		name      string
		audioFile string
		wantCode  int
		req       *request
	}

	tests := []test{
		{
			name:      "success",
			audioFile: filepath.Join("testdata", "cool-pic-128.jpg"),
			req: &request{
				method: http.MethodPut,
			},
			wantCode: http.StatusContinue,
		},
	}
	for _, tt := range tests {
		w := s.newRequest(endpoint, tt.req)

		b, err := jsonutil.ReadJson[model.FileUploadResponse](w.Body)
		s.Require().NoError(err)

		s.Equal(tt.wantCode, w.Code)
		s.NotEmpty(b.Location)
		s.NotEmpty(b.Links[0].Href)
		s.Equal("upload", b.Links[0].Rel)

	}
}
func (s *testSuite) TestUploadSongMetadata() {
	endpoint := "/v1/songs"

	type test struct {
		name string

		wantCode int
		req      *request
	}

	tests := []test{
		{
			name:     "no audio file href",
			wantCode: http.StatusBadRequest,
			req: &request{
				method: http.MethodPut,
				body: model.SongUploadRequest{
					Name:    "Cool music",
					Artists: []string{"rockguy", "rockguy2"},
					Genres:  []string{"Rock", "Pop"},
					// Audio:        "mysong.mp3",

					CreationDate: time.Now().Format(time.RFC3339),
				},
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			w := s.newRequest(endpoint, tt.req)
			s.Equal(tt.wantCode, w.Code)
			// s.Empty(w.Result().Header.Get("Location"))
		})
	}

	s.Run("success", func() {
		tt := test{
			req: &request{
				method: http.MethodPut,
				body: &model.SongUploadRequest{
					Name:         "Cool music",
					Artists:      []string{"rockguy", "rockguy2"},
					Genres:       []string{"Rock", "Pop"},
					Audio:        "mysong.mp3",
					Cover:        "asdasd",
					CreationDate: time.Now().Format(time.RFC3339),
				},
			},

			wantCode: http.StatusCreated,
		}
		w := s.newRequest(endpoint, tt.req)

		// err := json.NewEncoder(os.Stdout).Encode(tt.req.body)
		// t, err := json.Marshal(tt.req.body)
		// s.Equal("", string(t))
		// _, err := jsonutil.ReadJson[string](tt.req)
		// s.Require().NoError(err)
		s.Equal(tt.wantCode, w.Code)

		e2 := w.Header().Get("Location")
		s.NotEmpty(e2)

		// Check the uploaded song metadata
		w2 := s.newRequest("/v1/songs", &request{
			method: "GET",
		})

		s.Equal(http.StatusOK, w2.Code)

		b2, err := jsonutil.ReadJson[[]model.Song](w2.Body)
		s.Require().NoError(err, e2)
		s.NotNil(b2)

		// s.Equal("asdasd", b2)
	})
}

// TODO FINISH INTEGRATION TEST FOR FILE UPLOADING
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
