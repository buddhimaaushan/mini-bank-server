// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: verify_email.sql

package sqlc

import (
	"context"
)

const CreateVerifyEmail = `-- name: CreateVerifyEmail :one

INSERT INTO
    verify_emails (username, email, secret_code)
VALUES ($1, $2, $3) RETURNING id, username, email, secret_code, is_used, created_at, expired_at
`

type CreateVerifyEmailParams struct {
	Username   string `db:"username" json:"username"`
	Email      string `db:"email" json:"email"`
	SecretCode string `db:"secret_code" json:"secret_code"`
}

func (q *Queries) CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error) {
	row := q.db.QueryRow(ctx, CreateVerifyEmail, arg.Username, arg.Email, arg.SecretCode)
	var i VerifyEmail
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.SecretCode,
		&i.IsUsed,
		&i.CreatedAt,
		&i.ExpiredAt,
	)
	return i, err
}

const UpdateVerifyEmail = `-- name: UpdateVerifyEmail :one

UPDATE verify_emails
SET is_used = TRUE
WHERE
    id = $1
    AND secret_code = $2
    AND is_used = FALSE
    AND expired_at > now() RETURNING id, username, email, secret_code, is_used, created_at, expired_at
`

type UpdateVerifyEmailParams struct {
	ID         int64  `db:"id" json:"id"`
	SecretCode string `db:"secret_code" json:"secret_code"`
}

func (q *Queries) UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmail, error) {
	row := q.db.QueryRow(ctx, UpdateVerifyEmail, arg.ID, arg.SecretCode)
	var i VerifyEmail
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.SecretCode,
		&i.IsUsed,
		&i.CreatedAt,
		&i.ExpiredAt,
	)
	return i, err
}
