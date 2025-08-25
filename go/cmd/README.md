When using DBPoolDep found in altjoe/go-server-template/web/interfaces this is a generator that uses the following 
pointer functions

```
type DBPoolDep interface {
	GetDBPool() *pgxpool.Pool

	// here:
	Query(r context.Context, params DatabaseQueryParams, pool *pgxpool.Pool) (any, error)
	QueryTx(r context.Context, params DatabaseQueryParams, tx pgx.Tx) (any, error)
	// ^here

	GetTxWithCleanup(context.Context) (pgx.Tx, func(success bool), error)
}
```

type DatabaseQueryParams interface {
	Query(query *db.Queries, r context.Context) (any, error)
}
