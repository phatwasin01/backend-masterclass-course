package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()

}

type HashingTicketsResponse struct {
	UserID  int64  `json:"user_id"`
	EventID int64  `json:"event_id"`
	OrderID int64  `json:"order_id"`
	Hashed  string `json:"hashed"`
}

func (store *Store) CreateOrderTickets(ctx context.Context, arg CreateOrderParams) ([]Ticket, error) {
	var ticketlist []Ticket
	err := store.execTx(ctx, func(q *Queries) error {
		order, err := q.CreateOrder(ctx, arg)
		if err != nil {
			return err
		}
		for i := 0; i < int(order.Amount); i++ {
			ticket, err := q.CreateTicket(ctx, CreateTicketParams{
				UserID:     order.UserID,
				EventID:    order.EventID,
				OrderID:    order.ID,
				TicketUuid: uuid.New().String(),
			})
			ticketlist = append(ticketlist, ticket)
			if err != nil {
				return err
			}
		}
		if len(ticketlist) != int(order.Amount) {
			return errors.New("Ticket List != Order Amount")
		}
		err = q.UpdateEventSold(ctx, UpdateEventSoldParams{
			ID:         order.EventID,
			AmountSold: order.Amount,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return ticketlist, err
}
