-- name: CreateTransfer :one

INSERT INTO
    transfers (
        from_account_id,
        to_account_id,
        amount
    )
VALUES ($1, $2, $3) RETURNING *;

-- name: DeleteTransfer :exec

DELETE FROM transfers WHERE id = $1;

-- name: GetTransfer :one

SELECT * FROM transfers WHERE id = $1 LIMIT 1;

-- name: GetTransfersAsc :many

SELECT *
FROM transfers
WHERE
    from_account_id = $1
    AND to_account_id = $2
ORDER BY created_at ASC
LIMIT $3
OFFSET $4;

-- name: GetTransfersDesc :many

SELECT *
FROM transfers
WHERE
    from_account_id = $1
    AND to_account_id = $2
ORDER BY created_at Desc
LIMIT $3
OFFSET $4;

-- name: GetTransfersFromAsc :many

SELECT *
FROM transfers
WHERE from_account_id = $1
ORDER BY created_at ASC
LIMIT $2
OFFSET $3;

-- name: GetTransfersFromDesc :many

SELECT *
FROM transfers
WHERE from_account_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetTransfersToAsc :many

SELECT *
FROM transfers
WHERE to_account_id = $1
ORDER BY created_at ASC
LIMIT $2
OFFSET $3;

-- name: GetTransfersToDesc :many

SELECT *
FROM transfers
WHERE to_account_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetTransfersAmountEQAsc :many

SELECT *
FROM transfers
WHERE amount = $1
ORDER BY created_at ASC
LIMIT $2
OFFSET $3;

-- name: GetTransfersAmountEQDesc :many

SELECT *
FROM transfers
WHERE amount = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetTransfersAmountGTAsc :many

SELECT *
FROM transfers
WHERE amount > $1
ORDER BY created_at ASC
LIMIT $2
OFFSET $3;

-- name: GetTransfersAmountGTDesc :many

SELECT *
FROM transfers
WHERE amount > $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetTransfersAmountLTAsc :many

SELECT *
FROM transfers
WHERE amount < $1
ORDER BY created_at ASC
LIMIT $2
OFFSET $3;

-- name: GetTransfersAmountLTDesc :many

SELECT *
FROM transfers
WHERE amount < $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetTransfersAmountGTEQAsc :many

SELECT *
FROM transfers
WHERE amount >= $1
ORDER BY created_at ASC
LIMIT $2
OFFSET $3;

-- name: GetTransfersAmountGTEQDesc :many

SELECT *
FROM transfers
WHERE amount >= $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetTransfersAmountLTEQAsc :many

SELECT *
FROM transfers
WHERE amount <= $1
ORDER BY created_at ASC
LIMIT $2
OFFSET $3;

-- name: GetTransfersAmountLTEQDesc :many

SELECT *
FROM transfers
WHERE amount <= $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetTransfersAmountNOTEQAsc :many

SELECT *
FROM transfers
WHERE amount <> $1
ORDER BY created_at ASC
LIMIT $2
OFFSET $3;

-- name: GetTransfersAmountNOTEQDesc :many

SELECT *
FROM transfers
WHERE amount <> $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: GetTransfersAmountBetweenAsc :many

SELECT *
FROM transfers
WHERE amount BETWEEN $1 AND $2
ORDER BY created_at ASC
LIMIT $3
OFFSET $4;

-- name: GetTransfersAmountBetweenDesc :many

SELECT *
FROM transfers
WHERE amount BETWEEN $1 AND $2
ORDER BY created_at DESC
LIMIT $3
OFFSET $4;