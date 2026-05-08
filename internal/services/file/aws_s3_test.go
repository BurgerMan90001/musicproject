package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAWSS3(t *testing.T) {
	//t.Parallel()
	t.Skip("Skipping s3 test")
	ctx := t.Context()

	s3, err := NewS3(ctx, "us-east-1", "", "", "", "")
	require.NoError(t, err)

	t.Run("", func(t *testing.T) {
		content, err := s3.GetObject(ctx, "awdawda", "awdawdawd")

		require.NoError(t, err)
		assert.Equal(t, "", content)
	})
}
