package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/phatwasin01/ticketx-line-oa/util"
	"github.com/stretchr/testify/require"
)

func TestCreateOrder(t *testing.T) {
	//Create User
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
	//Create Organizer
	arg1 := CreateOrganizerParams{
		Name:     util.RandomUser(),
		Email:    util.RandomEmail(),
		Password: util.RandomString(4),
		Phone: sql.NullString{
			String: "0877896541",
			Valid:  true,
		},
	}
	orgainzer, err := testQueries.CreateOrganizer(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, orgainzer)

	//Create Event
	arg2 := CreateEventParams{
		Name:        util.RandomUser(),
		OrganizerID: orgainzer.ID,
		Price:       int32(util.RandomInt(10, 1000)),
		Amount:      int32(util.RandomInt(10, 1000)),
		StartTime:   time.Date(2023, time.April, 16, 0, 0, 0, 0, time.UTC),
	}
	event, err := testQueries.CreateEvent(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, event)
	store := NewStore(testDB)
	//Create Tickets
	amount := int32(util.RandomInt(1, 20))
	result, err := store.CreateOrderTickets(context.Background(), CreateOrderParams{
		UserID:   user.UserID,
		EventID:  event.ID,
		Amount:   amount,
		SumPrice: amount * event.Price,
	})
	fmt.Print(result)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
