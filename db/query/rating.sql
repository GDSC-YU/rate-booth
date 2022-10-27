-- name: AddRating :one
INSERT INTO rating (price, quality, booth_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListRatingsByBoothId :many
SELECT *
FROM rating
WHERE rating.booth_id = sqlc.arg(booth_id)
ORDER BY rating.created_at DESC;