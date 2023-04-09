-- name: CreateEvent :one
INSERT INTO events (
  name,
  organizer_id,
  price,
  amount,
  description,
  start_time
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetEvent :one
SELECT * FROM events
WHERE id = $1 LIMIT 1;

-- name: ListEvents :many
SELECT * FROM events
ORDER BY id
LIMIT $1 
OFFSET $2;

-- name: UpdateEventName :exec
UPDATE events 
SET name = $2
WHERE id = $1
RETURNING *;

-- name: UpdateEventPrice :exec
UPDATE events 
SET price = $2
WHERE id = $1
RETURNING *;

-- name: UpdateEventAmount :exec
UPDATE events 
SET amount = $2
WHERE id = $1
RETURNING *;

-- name: UpdateEventStart :exec
UPDATE events 
SET start_time = $2
WHERE id = $1
RETURNING *;

-- name: UpdateEventSold :exec
UPDATE events 
SET amount_sold = $2
WHERE id = $1
RETURNING *;

-- name: UpdateEventRedeem :exec
UPDATE events 
SET amount_redeem = $2
WHERE id = $1
RETURNING *;

-- name: DeleteEvent :exec
DELETE FROM events WHERE id = $1;