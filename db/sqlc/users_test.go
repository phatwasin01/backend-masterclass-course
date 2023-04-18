package db

import (
	"context"
	"testing"

	"github.com/phatwasin01/ticketx-line-oa/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		UserID:      util.RandomUser(),
		Email:       util.RandomEmail(),
		DisplayName: util.RandomUser(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.UserID, user.UserID)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.DisplayName, user.DisplayName)

	require.NotZero(t, user.UserID)
	require.NotZero(t, user.CreatedAt)
}
