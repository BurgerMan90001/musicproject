package encode

import "context"

type HLSEncoder interface {
	Segment(ctx context.Context, inputPath, outputDir string) error
}
