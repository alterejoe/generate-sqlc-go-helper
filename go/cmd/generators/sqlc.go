package generators

import (
	"fmt"

	"github.com/alterejoe/generate/sqlc-go-helper/cmd/interfaces"
)

func SqlcSelectGenerate(f interfaces.Sqlc) string {
	meta := fmt.Sprintf("-- name: %s%s %s\n", f.GetName(), f.GetIdentifier(), f.GetReturns())
	return fmt.Sprintf("%s%s FROM %s %s %s %s;",
		meta,
		f.GetQuery(),
		f.GetFrom(),
		f.GetWhere(),
		f.GetOrderBy(),
		f.GetLimit())
}
