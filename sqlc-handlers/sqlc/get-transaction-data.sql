-- name: GetLastUpdatedTransaction :one
SELECT
    last_updated
FROM
    "budget".transaction
ORDER BY
    last_updated DESC
LIMIT 1;

-- name: GetTransactions :many
SELECT
    id,
    posted,
    amount,
    description,
    payee,
    memo,
    transactedAt,
    pending
FROM
    "budget".transaction;

-- name: GetAccountTransactions :many
SELECT
    id,
    posted,
    amount,
    description,
    payee,
    memo,
    transactedAt,
    pending
FROM
    "budget".transaction
WHERE
    account_id = sqlc.arg (account_id);

