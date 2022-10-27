-- name: AddBooth :one
INSERT INTO booth (name, logo_url, description, twitter_url, instagram_url)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetBoothById :one
SELECT *
FROM booth
WHERE booth.id = sqlc.arg(booth_id);

-- name: ListBooths :many
SELECT *
FROM booth;