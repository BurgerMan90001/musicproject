// Uses commandline ffmpeg to encode files

package encode

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var _ HLSEncoder = (*FFmpeg)(nil)

type FFmpeg struct {
}

// Command line ffmpeg
func NewFFmpeg() *FFmpeg {
	return &FFmpeg{}
}

func (s *FFmpeg) Segment(ctx context.Context, inputPath, outputDir string) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("Create output dir: %w", err)
	}

	manifestPath := filepath.Join(outputDir, "index.m3u8")
	segmentPattern := filepath.Join(outputDir, "segment%03d.ts")

	// TODO
	cmd := exec.CommandContext(ctx,
		"ffmpeg",
		"-i", inputPath,
		"-codec:", "copy",
		// Segement duration in seconds
		"-hls_time", "6",
		// Keep all segments in manifest
		"-hls_list_size", "0",
		"-hls_segment_filename", segmentPattern,
		"-f", "hls",
		manifestPath,
	)

	//cmd.Stderr = os.Stderr

	// begin ffmpeg
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg exec: %w", err)
	}

	return nil
}
