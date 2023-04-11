-- name: CreateOrganizer :one
INSERT INTO organizers (
  name,
  email,
  password,
  phone
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetOrganizer :one
SELECT * FROM organizers
WHERE email = $1 LIMIT 1;

-- name: ListOrganizers :many
SELECT * FROM organizers
ORDER BY id
LIMIT $1 
OFFSET $2;

-- name: UpdateOrganizerPassword :exec
UPDATE organizers SET password = $2
WHERE id = $1
RETURNING *;

-- name: DeleteOrganizer :exec
DELETE FROM organizers WHERE id = $1;