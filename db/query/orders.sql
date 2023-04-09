-- name: CreateOrder :one
INSERT INTO orders (
  user_id,
  event_id,
  amount,
  sum_price,
  payment
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetOrderId :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;


-- name: ListOrdersEvent :many
SELECT * FROM orders
WHERE event_id = $1 ;

-- name: ListOrdersUser :many
SELECT * FROM orders
WHERE user_id = $1 ;


-- name: DeleteOrder :exec
DELETE FROM orders WHERE id = $1;