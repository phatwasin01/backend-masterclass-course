package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateEvent(t *testing.T) {
	arg := CreateEventParams{
		Name:        "YoungOhm",
		OrganizerID: 1,
		Price:       350,
		Amount:      100,
		StartTime:   time.Date(2023, time.April, 16, 0, 0, 0, 0, time.UTC),
	}

	user, err := testQueries.CreateEvent(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
}
