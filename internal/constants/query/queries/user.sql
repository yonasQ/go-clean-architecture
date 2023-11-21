-- name: CreateUser :one
INSERT INTO users (
    first_name,
    middle_name,
    last_name,
    email
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: UserByEmailExists :one
SELECT count_rows() FROM users where email = $1 AND deleted_at IS NULL; 

-- name: DeleteUser :one
UPDATE users set deleted_at =now() where id=$1 AND deleted_at IS NULL RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET 
  first_name = $1,
  middle_name = $2,
  last_name = $3
WHERE id = $4 ANd deleted_at IS NULL
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL;

-- name: GetUsers :many
SELECT * FROM users WHERE deleted_at IS NULL;
