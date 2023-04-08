-- name: CreateUser :one
INSERT INTO users (
  user_id,
  email,
  display_name
) VALUES (
  $1, $2, $3
)
RETURNING *;