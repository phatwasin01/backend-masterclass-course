-- name: CreateTicket :one
INSERT INTO tickets (
  user_id,
  event_id,
  order_id,
  hashed
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetTicketUser :one
SELECT * FROM tickets
WHERE user_id = $1;



-- name: RedeemTicket :exec
UPDATE Tickets SET is_redeemed = $2
WHERE id = $1
RETURNING *;

-- name: HashTicket :exec
UPDATE Tickets SET hashed = $2
WHERE id = $1
RETURNING *;


-- name: DeleteTicket :exec
DELETE FROM tickets WHERE id = $1;