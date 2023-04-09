package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateOrganizer(t *testing.T) {
	arg := CreateOrganizerParams{
		Name:     "G2",
		Email:    "g2@email.com",
		Password: "secret",
		Phone: sql.NullString{
			String: "0877896541",
			Valid:  true,
		},
	}

	user, err := testQueries.CreateOrganizer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
}
