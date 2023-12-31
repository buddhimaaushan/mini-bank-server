// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: account_holder.sql

package sqlc

import (
	"context"
)

const CreateAccountHolder = `-- name: CreateAccountHolder :one

INSERT INTO
    account_holders(acc_id, user_id)
VALUES ($1, $2)
RETURNING acc_id, user_id, created_at
`

type CreateAccountHolderParams struct {
	AccID  int64 `db:"acc_id" json:"acc_id"`
	UserID int64 `db:"user_id" json:"user_id"`
}

func (q *Queries) CreateAccountHolder(ctx context.Context, arg CreateAccountHolderParams) (AccountHolder, error) {
	row := q.db.QueryRow(ctx, CreateAccountHolder, arg.AccID, arg.UserID)
	var i AccountHolder
	err := row.Scan(&i.AccID, &i.UserID, &i.CreatedAt)
	return i, err
}

type CreateAccountHoldersParams struct {
	AccID  int64 `db:"acc_id" json:"acc_id"`
	UserID int64 `db:"user_id" json:"user_id"`
}

const DeleteAccountHolder = `-- name: DeleteAccountHolder :exec

DELETE FROM account_holders WHERE acc_id = $1 AND user_id = $2
`

type DeleteAccountHolderParams struct {
	AccID  int64 `db:"acc_id" json:"acc_id"`
	UserID int64 `db:"user_id" json:"user_id"`
}

func (q *Queries) DeleteAccountHolder(ctx context.Context, arg DeleteAccountHolderParams) error {
	_, err := q.db.Exec(ctx, DeleteAccountHolder, arg.AccID, arg.UserID)
	return err
}

const GetAccountHolder = `-- name: GetAccountHolder :one

SELECT acc_id, user_id, created_at FROM account_holders WHERE acc_id = $1 AND user_id = $2
`

type GetAccountHolderParams struct {
	AccID  int64 `db:"acc_id" json:"acc_id"`
	UserID int64 `db:"user_id" json:"user_id"`
}

func (q *Queries) GetAccountHolder(ctx context.Context, arg GetAccountHolderParams) (AccountHolder, error) {
	row := q.db.QueryRow(ctx, GetAccountHolder, arg.AccID, arg.UserID)
	var i AccountHolder
	err := row.Scan(&i.AccID, &i.UserID, &i.CreatedAt)
	return i, err
}

const GetAccountHoldersByAccountID = `-- name: GetAccountHoldersByAccountID :many

SELECT acc_id, user_id, created_at FROM account_holders WHERE acc_id = $1 LIMIT $2 OFFSET $3
`

type GetAccountHoldersByAccountIDParams struct {
	AccID  int64 `db:"acc_id" json:"acc_id"`
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

func (q *Queries) GetAccountHoldersByAccountID(ctx context.Context, arg GetAccountHoldersByAccountIDParams) ([]AccountHolder, error) {
	rows, err := q.db.Query(ctx, GetAccountHoldersByAccountID, arg.AccID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountHolder{}
	for rows.Next() {
		var i AccountHolder
		if err := rows.Scan(&i.AccID, &i.UserID, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetAccountHoldersByUserID = `-- name: GetAccountHoldersByUserID :many

SELECT acc_id, user_id, created_at FROM account_holders WHERE user_id = $1 LIMIT $2 OFFSET $3
`

type GetAccountHoldersByUserIDParams struct {
	UserID int64 `db:"user_id" json:"user_id"`
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

func (q *Queries) GetAccountHoldersByUserID(ctx context.Context, arg GetAccountHoldersByUserIDParams) ([]AccountHolder, error) {
	rows, err := q.db.Query(ctx, GetAccountHoldersByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountHolder{}
	for rows.Next() {
		var i AccountHolder
		if err := rows.Scan(&i.AccID, &i.UserID, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetAllAccountHoldersAsc = `-- name: GetAllAccountHoldersAsc :many

SELECT acc_id, user_id, created_at
FROM account_holders
ORDER BY
    created_at ASC,
    acc_id ASC,
    user_id ASC
LIMIT $1
OFFSET $2
`

type GetAllAccountHoldersAscParams struct {
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

func (q *Queries) GetAllAccountHoldersAsc(ctx context.Context, arg GetAllAccountHoldersAscParams) ([]AccountHolder, error) {
	rows, err := q.db.Query(ctx, GetAllAccountHoldersAsc, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountHolder{}
	for rows.Next() {
		var i AccountHolder
		if err := rows.Scan(&i.AccID, &i.UserID, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const GetAllAccountHoldersDesc = `-- name: GetAllAccountHoldersDesc :many

SELECT acc_id, user_id, created_at
FROM account_holders
ORDER BY
    created_at DESC,
    acc_id DESC,
    user_id DESC
LIMIT $1
OFFSET $2
`

type GetAllAccountHoldersDescParams struct {
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

func (q *Queries) GetAllAccountHoldersDesc(ctx context.Context, arg GetAllAccountHoldersDescParams) ([]AccountHolder, error) {
	rows, err := q.db.Query(ctx, GetAllAccountHoldersDesc, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountHolder{}
	for rows.Next() {
		var i AccountHolder
		if err := rows.Scan(&i.AccID, &i.UserID, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
