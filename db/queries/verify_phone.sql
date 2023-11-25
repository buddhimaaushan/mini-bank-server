-- name: CreateVerifyPhone :one

INSERT INTO
    verify_phones (username, phone, secret_code)
VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateVerifyPhone :one

UPDATE verify_phones
SET is_used = TRUE
WHERE
    id = @id
    AND secret_code = @secret_code
    AND is_used = FALSE
    AND expired_at > now() RETURNING *;