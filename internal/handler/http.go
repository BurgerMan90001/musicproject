package handler

import (
	"context"

	"songsled.com/pkg/model"
)

// func writeLocation(w http.ResponseWriter, v string) {
// 	w.Header().Set("Location", v)
// }

func contextClaims(ctx context.Context) (*model.Claims, bool) {
	claims, ok := ctx.Value("claims").(*model.Claims)
	return claims, ok
}
