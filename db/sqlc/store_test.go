package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateOrder(t *testing.T) {
	store := NewStore(testDB)

	result, err := store.CreateOrderTickets(context.Background(), CreateOrderParams{
		UserID:   int64(1),
		EventID:  int64(2),
		Amount:   int32(10),
		SumPrice: int32(3500),
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
