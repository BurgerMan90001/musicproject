package integration

func (s *testSuite) TestHandleSongMetadata() {
	url := "/v1/songs/"
	tests := []handlerTest{
		{
			Name: "success",
		},
		{
			Name: "",
		},
	}

	for _, tt := range tests {
		s.Run(tt.Name, func() {
			w := s.newRequest(url, tt.Req)

			s.Equal(tt.WantCode, w.Code)
		})
	}
}

func (s *testSuite) TestSearchSong() {

	tests := []handlerTest{
		{
			Name: "",
		},
	}

	for _, tt := range tests {
		s.Run(tt.Name, func() {

		})
	}

	s.Run("success", func() {

	})
}
