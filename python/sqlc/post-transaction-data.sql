-- name: BatchPostTransactionData :exec
INSERT INTO "budget".transaction (id, posted, amount, description, payee, memo, transactedAt, pending, account_id)
SELECT
    unnest(sqlc.arg (id)::text[]),
    unnest(sqlc.arg (posted)::int[]),
    unnest(sqlc.arg (amount)::text[]),
    unnest(sqlc.arg (description)::text[]),
    unnest(sqlc.arg (payee)::text[]),
    unnest(sqlc.arg (memo)::text[]),
    unnest(sqlc.arg (transactedAt)::int[]),
    unnest(sqlc.arg (pending)::bool[]),
    unnest(sqlc.arg (account_id)::text[])
ON CONFLICT (id)
    DO NOTHING;

