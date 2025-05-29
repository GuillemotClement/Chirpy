-- name: CreateChirps :one
INSERT INTO chirps (id, created_at, updated_at, body, user_id)
VALUEs (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1, 
  $2
)
RETURNING *;

