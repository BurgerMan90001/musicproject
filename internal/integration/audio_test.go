package integration

import "net/http"

func (s *testSuite) TestAudio() {
	s.T().Skip()

	url := "/v1/"

	// tempDir, err := os.MkdirTemp("", "audio")
	// s.Require().NoError(err)
	// s.T().Log(tempDir)

	// s.T().Cleanup(func() {
	// 	os.RemoveAll(tempDir)
	// })

	tests := []struct {
		name string
		file string

		//wantMessage string
		wantStatus int
	}{
		{
			name: "success",
			file: "8bitBossa.mp3",
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			req := &request{
				method: http.MethodGet,
			}
			w := s.newRequest(url+tt.file, req)

			// resBody, err := jsonutil.ReadJSON[map[string]any](w.Result().Body)
			// s.Require().NoError(err)

			//d, _ := io.ReadAll(w.Body)
			//s.Equal("", string(d))
			// s.Equal(tt.wantMessage, resBody["message"], tt.name)
			s.Equal(tt.wantStatus, w.Code, tt.name)
		})
	}
}
