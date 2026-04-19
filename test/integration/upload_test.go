package integration

import (
	"net/http"

	"musicproject.com/internal/jsonutil"
)

func (s *testSuite) TestUploadSongMetadata() {
	url := "/v1/upload/songs"

	tests := []HandlerTest{
		{
			Name: "empty request",
			Req: &request{
				method: http.MethodPost,
			},
			WantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		s.Run(tt.Name, func() {
			w := s.newRequest(url, tt.Req)
			s.Equal(tt.WantCode, w.Code)
		})
	}

	s.Run("success", func() {
		tt := HandlerTest{
			Req: &request{
				method: http.MethodPost,
				body: map[string]any{
					"genre":    "Rock",
					"filename": "mysong.mp3",
				},
			},
			WantCode: http.StatusOK,
		}
		w := s.newRequest(url, tt.Req)
		url := w.Result().Header.Get("Location")

		res := jsonutil.ReadJSONT[map[string]any](s.T(), w.Body)

		s.NotEmpty(url)
		s.Equal(tt.WantCode, w.Code, url)
		if res != nil {
			s.Equal("", res["message"])
		}
	})
}
