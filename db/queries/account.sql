-- name: UpdateAccountType :one

UPDATE accounts SET type = $2 WHERE id = $1 RETURNING *;

-- name: UpdateAccountAccStatus :one

UPDATE accounts SET acc_status = $2 WHERE id = $1 RETURNING *;

-- name: SetAccountBalance :one

UPDATE accounts SET balance = $2 WHERE id = $1 RETURNING *;

-- name: UpdateAccountBalance :one

UPDATE accounts
SET
    balance = balance + sqlc.arg(amount)
WHERE
    id = sqlc.arg(id) RETURNING *;

-- name: DeleteAccount :exec

DELETE FROM accounts WHERE id = $1;

-- name: CreateAccount :one

INSERT INTO
    accounts(type, balance, acc_status)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetAccounts :many

SELECT *
FROM accounts
LIMIT sqlc.arg(limit_no)
OFFSET sqlc.arg(offset_no);

-- name: GetAccount :one

SELECT * FROM accounts WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one

SELECT * FROM accounts WHERE id = $1 LIMIT 1 FOR NO KEY UPDATE;

-- name: GetAccountsByAccStatusAsc :many

SELECT *
FROM accounts
WHERE
    acc_status = sqlc.arg(active)
    OR acc_status = sqlc.arg(inactive)
    OR acc_status = sqlc.arg(holded)
    OR acc_status = sqlc.arg(deleted)
ORDER BY created_at ASC
LIMIT sqlc.arg(limit_no)
OFFSET sqlc.arg(offset_no);

-- name: GetAccountsByAccStatusDesc :many

SELECT *
FROM accounts
WHERE
    acc_status = sqlc.arg(active)
    OR acc_status = sqlc.arg(inactive)
    OR acc_status = sqlc.arg(holded)
    OR acc_status = sqlc.arg(deleted)
ORDER BY created_at DESC
LIMIT sqlc.arg(limit_no)
OFFSET sqlc.arg(offset_no);

-- name: GetAccountsBalanceEQAsc :many

SELECT *
FROM accounts
WHERE
    balance = $1
    AND (
        acc_status = sqlc.arg(active)
        OR acc_status = sqlc.arg(inactive)
        OR acc_status = sqlc.arg(holded)
        OR acc_status = sqlc.arg(deleted)
    )
ORDER BY created_at ASC
LIMIT sqlc.arg(limit_no)
OFFSET sqlc.arg(offset_no);

-- name: GetAccountsBalanceEQDesc :many

SELECT *
FROM accounts
WHERE
    balance = $1
    AND (
        acc_status = sqlc.arg(active)
        OR acc_status = sqlc.arg(inactive)
        OR acc_status = sqlc.arg(holded)
        OR acc_status = sqlc.arg(deleted)
    )
ORDER BY created_at DESC
LIMIT sqlc.arg(limit_no)
OFFSET sqlc.arg(offset_no);

-- name: GetAccountsBalanceGTAsc :many

SELECT *
FROM accounts
WHERE
    balance > $1
    AND (
        acc_status = sqlc.arg(active)
        OR acc_status = sqlc.arg(inactive)
        OR acc_status = sqlc.arg(holded)
        OR acc_status = sqlc.arg(deleted)
    )
ORDER BY created_at ASC
LIMIT sqlc.arg(limit_no)
OFFSET sqlc.arg(offset_no);

-- name: GetAccountsBalanceGTDesc :many

SELECT *
FROM accounts
WHERE
    balance > $1
    AND (
        acc_status = sqlc.arg(active)
        OR acc_status = sqlc.arg(inactive)
        OR acc_status = sqlc.arg(holded)
        OR acc_status = sqlc.arg(deleted)
    )
ORDER BY created_at DESC
LIMIT sqlc.arg(limit_no)
OFFSET sqlc.arg(offset_no);

-- name: GetAccountsBalanceLTAsc :many

SELECT *
FROM accounts
WHERE
    balance < $1
    AND (
        acc_status = sqlc.arg(active)
        OR acc_status = sqlc.arg(inactive)
        OR acc_status = sqlc.arg(holded)
        OR acc_status = sqlc.arg(deleted)
    )
ORDER BY created_at ASC
LIMIT sqlc.arg(limit_no)
OFFSET sqlc.arg(offset_no);

-- name: GetAccountsBalanceLTDesc :many

SELECT *
FROM accounts
WHERE
    balance < $1
    AND (
        acc_status = sqlc.arg(active)
        OR acc_status = sqlc.arg(inactive)
        OR acc_status = sqlc.arg(holded)
        OR acc_status = sqlc.arg(deleted)
    )
ORDER BY created_at DESC
LIMIT sqlc.arg(limit_no)
OFFSET sqlc.arg(offset_no);

-- name: GetAccountsBalanceGTEQAsc :many

SELECT *
FROM accounts
WHERE
    balance >= $1
    AND (
        acc_status = sqlc.arg(active)
        OR acc_status = sqlc.arg(inactive)
        OR acc_status = sqlc.arg(holded)
        OR acc_status = sqlc.arg(deleted)
    )
ORDER BY created_at ASC
LIMIT sqlc.arg(limit_no)
OFFSET sqlc.arg(offset_no);

-- name: GetAccountsBalanceGTEQDesc :many

SELECT *
FROM accounts
WHERE
    balance >= $1
    AND (
        acc_status = sqlc.arg(active)
        OR acc_status = sqlc.arg(inactive)
        OR acc_status = sqlc.arg(holded)
        OR acc_status = sqlc.arg(deleted)
    )
ORDER BY created_at DESC
LIMIT sqlc.arg(limit_no)
OFFSET sqlc.arg(offset_no);

-- name: GetAccountsBalanceLTEQAsc :many

SELECT *
FROM accounts
WHERE
    balance <= $1
    AND (
        acc_status = sqlc.arg(active)
        OR acc_status = sqlc.arg(inactive)
        OR acc_status = sqlc.arg(holded)
        OR acc_status = sqlc.arg(deleted)
    )
ORDER BY created_at ASC
LIMIT sqlc.arg(limit_no)
OFFSET sqlc.arg(offset_no);

-- name: GetAccountsBalanceLTEQDesc :many

SELECT *
FROM accounts
WHERE
    balance <= $1
    AND (
        acc_status = sqlc.arg(active)
        OR acc_status = sqlc.arg(inactive)
        OR acc_status = sqlc.arg(holded)
        OR acc_status = sqlc.arg(deleted)
    )
ORDER BY created_at DESC
LIMIT sqlc.arg(limit_no)
OFFSET sqlc.arg(offset_no);

-- name: GetAccountsBalanceNOTEQAsc :many

SELECT *
FROM accounts
WHERE
    balance <> $1
    AND (
        acc_status = sqlc.arg(active)
        OR acc_status = sqlc.arg(inactive)
        OR acc_status = sqlc.arg(holded)
        OR acc_status = sqlc.arg(deleted)
    )
ORDER BY created_at ASC
LIMIT sqlc.arg(limit_no)
OFFSET sqlc.arg(offset_no);

-- name: GetAccountsBalanceNOTEQDesc :many

SELECT *
FROM accounts
WHERE
    balance <> $1
    AND (
        acc_status = sqlc.arg(active)
        OR acc_status = sqlc.arg(inactive)
        OR acc_status = sqlc.arg(holded)
        OR acc_status = sqlc.arg(deleted)
    )
ORDER BY created_at DESC
LIMIT sqlc.arg(limit_no)
OFFSET sqlc.arg(offset_no);

-- name: GetAccountsBalanceBetweenAsc :many

SELECT *
FROM accounts
WHERE (
        acc_status = sqlc.arg(active)
        OR acc_status = sqlc.arg(inactive)
        OR acc_status = sqlc.arg(holded)
        OR acc_status = sqlc.arg(deleted)
    )
    AND balance BETWEEN $1 AND $2
ORDER BY created_at ASC
LIMIT sqlc.arg(limit_no)
OFFSET sqlc.arg(offset_no);

-- name: GetAccountsBalanceBetweenDesc :many

SELECT *
FROM accounts
WHERE (
        acc_status = sqlc.arg(active)
        OR acc_status = sqlc.arg(inactive)
        OR acc_status = sqlc.arg(holded)
        OR acc_status = sqlc.arg(deleted)
    )
    AND balance BETWEEN $1 AND $2
ORDER BY created_at DESC
LIMIT sqlc.arg(limit_no)
OFFSET sqlc.arg(offset_no);