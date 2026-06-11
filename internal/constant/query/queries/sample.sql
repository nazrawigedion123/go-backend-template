-- name: CreateSample :one
INSERT INTO sample (
    name, 
    email, 
    age, 
    is_active
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetSampleByID :one
SELECT * FROM sample 
WHERE id = $1 
LIMIT 1;

-- name: GetSampleByEmail :one
SELECT * FROM sample 
WHERE email = $1 
LIMIT 1;

-- name: GetAllSamples :many
SELECT * FROM sample 
ORDER BY created_at DESC;

-- name: GetAllActiveSamples :many
SELECT * FROM sample 
WHERE is_active = true 
ORDER BY created_at DESC;

-- name: UpdateSample :one
UPDATE sample 
SET 
    name = COALESCE(sqlc.narg('name'), name),
    email = COALESCE(sqlc.narg('email'), email),
    age = COALESCE(sqlc.narg('age'), age),
    is_active = COALESCE(sqlc.narg('is_active'), is_active),
    updated_at = CURRENT_TIMESTAMP
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteSample :exec
DELETE FROM sample 
WHERE id = $1;

-- name: DeleteSampleByEmail :exec
DELETE FROM sample 
WHERE email = $1;

-- name: GetSamplesPaginated :many
SELECT * FROM sample 
ORDER BY created_at DESC 
LIMIT $1 OFFSET $2;

-- name: CountSamples :one
SELECT COUNT(*) FROM sample;

-- name: CountActiveSamples :one
SELECT COUNT(*) FROM sample 
WHERE is_active = true;