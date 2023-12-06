-- name: CreateUser :one

INSERT INTO
    users (
        first_name,
        last_name,
        username,
        nic,
        hashed_password,
        email,
        phone
    )
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: DeleteUser :one

DELETE FROM users WHERE id = $1:: BIGINT RETURNING *;

-- name: UpdateUser :one

UPDATE users
SET
    first_name = COALESCE(
        sqlc.narg(first_name),
        first_name
    ),
    last_name = COALESCE(
        sqlc.narg(last_name),
        last_name
    ),
    hashed_password = COALESCE(
        sqlc.narg(hashed_password),
        hashed_password
    ),
    password_changed_at = COALESCE(
        sqlc.narg(password_changed_at),
        password_changed_at
    ),
    email = COALESCE(sqlc.narg(email), email),
    is_email_verified = COALESCE(
        sqlc.narg(is_email_verified),
        is_email_verified
    ),
    email_changed_at = COALESCE(
        sqlc.narg(email_changed_at),
        email_changed_at
    ),
    email = COALESCE(sqlc.narg(email), email),
    is_email_verified = COALESCE(
        sqlc.narg(is_email_verified),
        is_email_verified
    ),
    email_changed_at = COALESCE(
        sqlc.narg(email_changed_at),
        email_changed_at
    ),
    phone = COALESCE(sqlc.narg(phone), phone),
    is_phone_verified = COALESCE(
        sqlc.narg(is_phone_verified),
        is_phone_verified
    ),
    phone_changed_at = COALESCE(
        sqlc.narg(phone_changed_at),
        phone_changed_at
    ),
    acc_status = COALESCE(
        sqlc.narg(acc_status),
        acc_status
    ),
    customer_rank = COALESCE(
        sqlc.narg(customer_rank),
        customer_rank
    ),
    is_an_employee = COALESCE(
        sqlc.narg(is_an_employee),
        is_an_employee
    ),
    is_a_customer = COALESCE(
        sqlc.narg(is_a_customer),
        is_a_customer
    ),
    role = COALESCE(sqlc.narg(role), role),
    department = COALESCE(
        sqlc.narg(department),
        department
    )
WHERE
    id = sqlc.arg(id) RETURNING *;

-- name: GetAllUsersOrderByIDAsc :many

SELECT * FROM users ORDER BY id ASC LIMIT $1 OFFSET $2;

-- name: GetAllUsersOrderByIDDesc :many

SELECT * FROM users ORDER BY id DESC LIMIT $1 OFFSET $2;

-- name: GetAllUsersAsc :many

SELECT *
FROM users
ORDER BY
    first_name ASC,
    last_name ASC,
    username ASC
LIMIT $1
OFFSET $2;

-- name: GetAllUsersDesc :many

SELECT *
FROM users
ORDER BY
    first_name ASC,
    last_name ASC,
    username ASC
LIMIT $1
OFFSET $2;

-- name: GetUsersByPatternAsc :many

SELECT *
FROM users
WHERE
    first_name LIKE sqlc.arg(str_pattern):: text
    OR last_name LIKE sqlc.arg(str_pattern):: text
    OR username LIKE sqlc.arg(str_pattern):: text
    OR nic LIKE sqlc.arg(str_pattern):: text
    OR email LIKE sqlc.arg(str_pattern):: text
    OR phone LIKE sqlc.arg(str_pattern):: text
    OR role LIKE sqlc.arg(str_pattern):: text
    OR department LIKE sqlc.arg(str_pattern):: text
ORDER BY
    first_name ASC,
    last_name ASC,
    username ASC
LIMIT $1
OFFSET $2;

-- name: GetUsersByPatternDesc :many

SELECT *
FROM users
WHERE
    first_name LIKE sqlc.arg(str_pattern):: text
    OR last_name LIKE sqlc.arg(str_pattern):: text
    OR username LIKE sqlc.arg(str_pattern):: text
    OR nic LIKE sqlc.arg(str_pattern):: text
    OR email LIKE sqlc.arg(str_pattern):: text
    OR phone LIKE sqlc.arg(str_pattern):: text
    OR role LIKE sqlc.arg(str_pattern):: text
    OR department LIKE sqlc.arg(str_pattern):: text
ORDER BY
    first_name DESC,
    last_name DESC,
    username DESC
LIMIT $1
OFFSET $2;

-- name: GetUsersByPatternAndAccStatusAsc :many

SELECT *
FROM users
WHERE (
        first_name LIKE sqlc.arg(str_pattern):: text
        OR last_name LIKE sqlc.arg(str_pattern):: text
        OR username LIKE sqlc.arg(str_pattern):: text
        OR nic LIKE sqlc.arg(str_pattern):: text
        OR email LIKE sqlc.arg(str_pattern):: text
        OR phone LIKE sqlc.arg(str_pattern):: text
        OR role LIKE sqlc.arg(str_pattern):: text
        OR department LIKE sqlc.arg(str_pattern):: text
    )
    AND (
        acc_status = sqlc.arg(status_active)
        OR acc_status = sqlc.arg(status_inactive)
        OR acc_status = sqlc.arg(status_deleted)
        OR acc_status = sqlc.arg(status_holded)
    )
ORDER BY
    first_name ASC,
    last_name ASC,
    username ASC
LIMIT $1
OFFSET $2;

-- name: GetUsersByPatternAndAccStatusDesc :many

SELECT *
FROM users
WHERE (
        first_name LIKE sqlc.arg(str_pattern):: text
        OR last_name LIKE sqlc.arg(str_pattern):: text
        OR username LIKE sqlc.arg(str_pattern):: text
        OR nic LIKE sqlc.arg(str_pattern):: text
        OR email LIKE sqlc.arg(str_pattern):: text
        OR phone LIKE sqlc.arg(str_pattern):: text
        OR role LIKE sqlc.arg(str_pattern):: text
        OR department LIKE sqlc.arg(str_pattern):: text
    )
    AND (
        acc_status = sqlc.arg(status_active)
        OR acc_status = sqlc.arg(status_inactive)
        OR acc_status = sqlc.arg(status_deleted)
        OR acc_status = sqlc.arg(status_holded)
    )
ORDER BY
    first_name DESC,
    last_name DESC,
    username DESC
LIMIT $1
OFFSET $2;

-- name: GetUsersByAccStatusAsc :many

SELECT *
FROM users
WHERE acc_status = $1
ORDER BY
    first_name ASC,
    last_name ASC,
    username ASC
LIMIT $1
OFFSET $2;

-- name: GetUsersByAccStatusDesc :many

SELECT *
FROM users
WHERE acc_status = $1
ORDER BY
    first_name DESC,
    last_name DESC,
    username DESC
LIMIT $1
OFFSET $2;

-- name: GetUsersByCustRankAsc :many

SELECT *
FROM users
WHERE customer_rank = $1
ORDER BY
    first_name ASC,
    last_name ASC,
    username ASC
LIMIT $1
OFFSET $2;

-- name: GetUsersByCustRankDesc :many

SELECT *
FROM users
WHERE customer_rank = $1
ORDER BY
    first_name DESC,
    last_name DESC,
    username DESC
LIMIT $1
OFFSET $2;

-- name: GetUserByID :one

SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one

SELECT * FROM users WHERE username = $1 LIMIT 1;