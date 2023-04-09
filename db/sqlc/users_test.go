package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		UserID:      "U4af4980629",
		Email:       "lineOA@email.com",
		DisplayName: "LineOA",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.UserID, user.UserID)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.DisplayName, user.DisplayName)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
}
