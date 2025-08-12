package dto

import (
	"testing"

	"github.com/Racuwcka/shorter-url/internal/handler/shortener"
	"github.com/stretchr/testify/require"
)

func TestLinkRequestValidate(t *testing.T) {
	in := LinkRequest{
		Link: "https://google.com",
	}

	err := in.Validate()
	require.NoError(t, err)
}

func TestValidateError(t *testing.T) {
	cases := []struct {
		name    string
		request LinkRequest
		expErr  error
	}{
		{
			name: "link is empty",
			request: LinkRequest{
				Link: "",
			},
			expErr: shortener.ErrEmptyLink,
		},
		{
			name: "link is not url",
			request: LinkRequest{
				Link: "hello",
			},
			expErr: shortener.ErrInvalidURL,
		},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			err := tCase.request.Validate()
			require.Error(t, err)
			require.EqualError(t, tCase.expErr, err.Error())
		})
	}
}
