package streaming

type HLS struct {
}

func NewHLS() *HLS {
	return &HLS{}
}

func (s *HLS) Stream(file string, segmentDuration int) error {
	return nil
}
