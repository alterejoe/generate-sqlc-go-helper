-- name: GetLastUpdated :one
SELECT
    last_updated
FROM
    "budget".account
ORDER BY
    last_updated DESC
LIMIT 1;

-- name: GetAccounts :many
SELECT
    id,
    name,
    balance,
    org_id,
    org_name,
    last_updated
FROM
    "budget".account;

-- name: GetFirstAccount :one
SELECT
    id,
    name
FROM
    "budget".account
LIMIT 1;

