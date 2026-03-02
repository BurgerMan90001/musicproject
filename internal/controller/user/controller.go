package user

import (
	"context"
	"errors"
	"regexp"

	"okapi.com/internal/auth"
	"okapi.com/internal/controller"
	"okapi.com/internal/repository"
	"okapi.com/pkg/model"
)

//var ErrInvalidEFormat

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

func (c *Controller) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := c.repo.GetUserByEmail(ctx, email)

	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, repository.ErrNotFound
	}
	return user, nil
}

func (c *Controller) PutUser(ctx context.Context, id string, email string, password string) error {

	user := &model.User{
		ID:           id,
		Email:        email,
		PasswordHash: auth.HashPassword(password),
	}
	valid, err := isValidEmailString(email)

	if !valid {
		return controller.ErrInvalidFormat
	}
	if err != nil {
		return err
	}

	if err := c.repo.PutUser(ctx, user.ID, user); err != nil {
		return err
	}

	return nil
}

func (c *Controller) DeleteUserByID(ctx context.Context, id string) error {
	_, err := c.repo.GetUserByID(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return repository.ErrNotFound
	}
	if err := c.repo.DeleteUserByID(ctx, id); err != nil {
		return err
	}
	return nil
}
func isValidEmailString(email string) (bool, error) {
	return regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email)
}
