-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email)
VALUEs (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1
)
RETURNING *;

-- name: DeleteAllUser :exec
DELETE FROM users;