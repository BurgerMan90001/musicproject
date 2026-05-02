package integration

func (s *testSuite) TestHandleSongMetadata() {
	s.T().Skip()
	url := "/v1/songs/"
	tests := []handlerTest{
		{
			Name: "success",
			Req:  &request{},
		},
		{
			Name: "",
			Req:  &request{},
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
