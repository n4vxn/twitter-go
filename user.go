package twitter

import (
	"context"
	"errors"
	"time"
)

var (
	ErrUsernameTaken = errors.New("username taken")
	ErrEmailTaken = errors.New("email taken")
)

type UserRepo interface {
	Create(context.Context, User) (User, error)
	GetByUsername(context.Context, string) (User, error)
	GetByEmail(context.Context, string) (User, error)
}

type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
