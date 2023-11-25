-- name: CreateAccountHolder :one

INSERT INTO
    account_holders(acc_id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: CreateAccountHolders :copyfrom

INSERT INTO account_holders(acc_id, user_id) VALUES ($1, $2);

-- name: DeleteAccountHolder :exec

DELETE FROM account_holders WHERE acc_id = $1 AND user_id = $2;

-- name: GetAccountHolder :one

SELECT * FROM account_holders WHERE acc_id = $1 AND user_id = $2;

-- name: GetAccountHoldersByAccountID :many

SELECT * FROM account_holders WHERE acc_id = $1 LIMIT $2 OFFSET $3;

-- name: GetAccountHoldersByUserID :many

SELECT * FROM account_holders WHERE user_id = $1 LIMIT $2 OFFSET $3;

-- name: GetAllAccountHoldersAsc :many

SELECT *
FROM account_holders
ORDER BY
    created_at ASC,
    acc_id ASC,
    user_id ASC
LIMIT $1
OFFSET $2;

-- name: GetAllAccountHoldersDesc :many

SELECT *
FROM account_holders
ORDER BY
    created_at DESC,
    acc_id DESC,
    user_id DESC
LIMIT $1
OFFSET $2;