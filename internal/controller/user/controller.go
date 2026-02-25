package user

import (
	"context"
	"errors"

	"movieexample.com/internal/repository"
	"movieexample.com/pkg/model"
)

type Controller struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user, err := c.repo.GetUserByID(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, repository.ErrNotFound
	}
	return user, nil
}
func (c *Controller) PutUser(ctx context.Context, u *model.User) error {
	return c.repo.PutUser(ctx, u.ID, u)
}
