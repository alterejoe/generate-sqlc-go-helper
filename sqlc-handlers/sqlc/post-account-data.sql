-- name: PostAccountData :exec
INSERT INTO "budget".account (id, name, balance, org_id, org_name, last_updated)
    VALUES (sqlc.arg (id), sqlc.arg (name), sqlc.arg (balance), sqlc.arg (org_id), sqlc.arg (org_name), sqlc.arg (last_updated));

-- name: BatchPostAccountData :exec
INSERT INTO "budget".account (id, name, balance, org_id, org_name)
SELECT
    unnest(sqlc.arg (id)::text[]),
    unnest(sqlc.arg (name)::text[]),
    unnest(sqlc.arg (balance)::float[]),
    unnest(sqlc.arg (org_id)::text[]),
    unnest(sqlc.arg (org_name)::text[])
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    balance = EXCLUDED.balance,
    org_id = EXCLUDED.org_id,
    org_name = EXCLUDED.org_name;

