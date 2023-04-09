package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/phatwasin01/ticketx-line-oa/util"
	"github.com/stretchr/testify/require"
)

func TestCreateOrganizer(t *testing.T) {
	arg := CreateOrganizerParams{
		Name:     util.RandomUser(),
		Email:    util.RandomEmail(),
		Password: util.RandomString(4),
		Phone: sql.NullString{
			String: "0877896541",
			Valid:  true,
		},
	}

	user, err := testQueries.CreateOrganizer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
}
