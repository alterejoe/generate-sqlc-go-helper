package generators

import (
	"fmt"

	"github.com/alterejoe/generate/sqlc-go-helper/cmd/interfaces"
)

func SqlcSelectGenerate(f interfaces.DisplaySqlc) string {
	meta := fmt.Sprintf("-- name: %s%s %s\n", f.GetName(), f.GetIdentifier(), f.GetReturns())
	return fmt.Sprintf("%sSELECT %s FROM %s %s %s %s;",
		meta,
		f.GetSelect(),
		f.GetFrom(),
		f.GetWhere(),
		f.GetOrderBy(),
		f.GetLimit())
}
