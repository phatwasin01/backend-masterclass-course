package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/phatwasin01/ticketx-line-oa/util"
	"github.com/stretchr/testify/require"
)

func TestCreateEvent(t *testing.T) {
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

	arg1 := CreateEventParams{
		Name:        util.RandomUser(),
		OrganizerID: user.ID,
		Price:       int32(util.RandomInt(10, 1000)),
		Amount:      int32(util.RandomInt(10, 1000)),
		StartTime:   time.Date(2023, time.April, 16, 0, 0, 0, 0, time.UTC),
	}
	user1, err := testQueries.CreateEvent(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, user1)
}
