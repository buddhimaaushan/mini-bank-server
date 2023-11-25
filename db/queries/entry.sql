-- name: CreateEntry :one

INSERT INTO
    entries (account_id, amount)
VALUES ($1, $2) RETURNING *;

-- name: DeleteEntry :exec

DELETE FROM entries WHERE id = $1;

-- name: GetEntry :one

SELECT * FROM entries WHERE id = $1 LIMIT 1;

-- name: GetEntriesAsc :many

SELECT *
FROM entries
WHERE account_id = $1
ORDER BY
    created_at ASC,
    account_id ASC
LIMIT $2
OFFSET $3;

-- name: GetEntriesDesc :many

SELECT *
FROM entries
WHERE account_id = $1
ORDER BY
    created_at DESC,
    account_id DESC
LIMIT $2
OFFSET $3;