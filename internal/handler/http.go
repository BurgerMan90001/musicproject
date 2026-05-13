package handler

import (
	"context"
	"net/url"
	"os"

	"songsled.com/pkg/model"
)

func contextClaims(ctx context.Context) (*model.Claims, bool) {
	claims, ok := ctx.Value("claims").(*model.Claims)
	return claims, ok
}

// Start with the version number (ex: v1).
// Then the rest of the path.
func apiJoinUrl(e ...string) (string, error) {
	return url.JoinPath(os.Getenv("API_URL"), e...)
}
