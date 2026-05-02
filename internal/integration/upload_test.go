package integration

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
)

func (s *testSuite) TestUploadSongMetadata() {
	endpoint := "/v1/songs/upload"

	token, err := s.jwtAccess.GenerateToken(uuid.Nil)
	s.Require().NoError(err)

	tests := []handlerTest{
		{
			Name: "empty request",
			Req: &request{
				method:      http.MethodPost,
				accessToken: token,
			},
			WantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		s.Run(tt.Name, func() {
			w := s.newRequest(endpoint, tt.Req)
			s.Equal(tt.WantCode, w.Code)
		})
	}

	s.Run("success", func() {
		tt := handlerTest{
			Req: &request{
				method: http.MethodPost,
				body: map[string]any{
					"name":     "Cool music",
					"genre":    "Rock",
					"filename": "mysong.mp3",
				},
				accessToken: token,
			},
			WantCode: http.StatusOK,
		}
		w := s.newRequest(endpoint, tt.Req)
		location := w.Result().Header.Get("Location")

		s.NotEmpty(location)
		s.Equal(tt.WantCode, w.Code, location)

		s.Contains(location, tt.Req.body["filename"])

		if !testing.Short() && os.Getenv("USE_CLOUD") == "true" {
			f, err := os.Open(filepath.Join("testdata", "song18.mp3"))
			s.Require().NoError(err)

			r, err := http.NewRequest("PUT", location, f)
			s.Require().NoError(err)

			res, err := http.DefaultClient.Do(r)
			s.Require().NoError(err)

			data, err := io.ReadAll(res.Body)
			s.Require().NoError(err)

			s.Equal("", string(data))
		}

		// Check the uploaded song metadata
		metadataEndpoint, err := url.JoinPath("/v1/songs")
		s.Require().NoError(err)

		w2 := s.newRequest(metadataEndpoint, &request{
			method: "GET",
		})

		s.Equal(http.StatusOK, w2.Code)

	})
}
