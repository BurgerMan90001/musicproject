package handler

import (
	"context"

	"musicproject.com/pkg/model"
)

func contextClaims(ctx context.Context) (*model.Claims, bool) {
	claims, ok := ctx.Value("claims").(*model.Claims)
	return claims, ok
}
