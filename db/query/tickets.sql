-- name: CreateTicket :one
INSERT INTO tickets (
  user_id,
  event_id,
  order_id,
  ticket_uuid
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetTicketUser :one
SELECT * FROM tickets
WHERE user_id = $1;

-- name: GetTicketOrder :many
SELECT * FROM tickets
WHERE order_id = $1 AND user_id = $2;


-- name: RedeemTicket :exec
UPDATE tickets SET is_redeemed = $2
WHERE id = $1
RETURNING (user_id,event_id,order_id,ticket_uuid);



-- name: DeleteTicket :exec
DELETE FROM tickets WHERE id = $1;