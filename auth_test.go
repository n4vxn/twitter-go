package twitter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegisterInput_Validate(t *testing.T) {
	testCases := []struct {
		name  string
		input RegisterInput
		err   error
	}{
		{
			name: "valid",
			input: RegisterInput{
				Username:        "bob",
				Email:           "bob@gmail.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: nil,
		}, {
			name: "invalid email",
			input: RegisterInput{
				Username:        "bob",
				Email:           "bob",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: ErrValidation,
		}, {
			name: "too short username",
			input: RegisterInput{
				Username:        "b",
				Email:           "bob",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: ErrValidation,
		}, {
			name: "too short password",
			input: RegisterInput{
				Username:        "bob",
				Email:           "bob",
				Password:        "pass",
				ConfirmPassword: "password",
			},
			err: ErrValidation,
		}, {
			name: "confirm password doesnt match password",
			input: RegisterInput{
				Username:        "bob",
				Email:           "bob@gmail.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: ErrValidation,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.Validate()
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
