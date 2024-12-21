package twitter

import (
	"context"
	"fmt"
	"regexp"
	"strings"

)

var (
	UsernameMinLength = 2
	PasswordMinLength = 6
)

type AuthService interface {
	Register(context.Context, RegisterInput) (AuthResponse, error)
	Login(context.Context, LoginInput) (AuthResponse, error)
}

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+[a-zA-Z]{2,}$")

type AuthResponse struct {
	AccessToken string
	User        User
}

type RegisterInput struct {
	Email           string
	Username        string
	Password        string
	ConfirmPassword string
}

type LoginInput struct {
	Email           string
	Password        string
}

func (in *RegisterInput) Sanitize() {
	in.Email = strings.TrimSpace(in.Email)
	in.Email = strings.ToLower(in.Email)

	in.Username = strings.TrimSpace(in.Username)
}

func (in *RegisterInput) Validate() error {

	if len(in.Username) < UsernameMinLength {
		return fmt.Errorf("%w: username not long enough, (%d) character atleast", ErrValidation, UsernameMinLength)
	}

	if !emailRegexp.MatchString(in.Email) {
		return fmt.Errorf("%w: invalid email format", ErrValidation)
	}

	if len(in.Password) < PasswordMinLength {
		return fmt.Errorf("%w: password must be at least %d characters long", ErrValidation, PasswordMinLength)
	}

	if in.Password != in.ConfirmPassword {
		return fmt.Errorf("%w: passwords do not match", ErrValidation)
	}

	return nil
}

func (in *LoginInput) Sanitize() {
	in.Email = strings.TrimSpace(in.Email)
	in.Email = strings.ToLower(in.Email)
}

func (in *LoginInput) Validate() error {
	emailRegexp := regexp.MustCompile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+[a-zA-Z]{2,}$")

	if !emailRegexp.MatchString(in.Email) {
		return fmt.Errorf("%w: invalid email format", ErrValidation)
	}

	if len(in.Password) < 1 {
		return fmt.Errorf("%w: password required", ErrValidation)
	}

	return nil
}
