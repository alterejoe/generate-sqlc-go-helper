package interfaces

import (
	"log/slog"

	"github.com/dave/dst"
)

type Standard interface {
	GetName() string
	GetLowerName() string
	GetAbbv() string
}

// these will be duplicated from the origin Node
// ex: GenStruct will be a struct data created from a GenDecl
type Struct interface {
	Standard
	// THIS IS WHERE TYOU WILL PUT FUNCITONS THAT THE DISPLAY
	// STRUCT WILL USE
	GetStructFields() []*dst.Field
	GetLogger() *slog.Logger
}

type Function interface {
	Standard
	// THIS IS WHERE TYOU WILL PUT FUNCITONS THAT THE DISPLAY
	// FUNCTION WILL USE
	GetParams() []*dst.Field
	GetReturns() []*dst.Field
}

type DisplayStruct interface {
	Struct // to provide data to the template
	GetTypeSpec() *dst.TypeSpec
}

type DisplayFunction interface {
	Function // to provide data to the template
	GetGenerateFunctionName() string
	GetGenerateReceiver() *dst.FieldList
	GetGenerateFunctionParams() *dst.FieldList
	GetGenerateResults() *dst.FieldList
	GetBody() *dst.BlockStmt
}

type SqlcSelect interface {
	Standard
	GetSelect() string
	GetIdentifier() string
	GetFrom() string
	GetWhere() string
	GetOrderBy() string
	GetLimit() string
	GetReturns() string
}

type SqlcInsert interface {
	GetInsert() string
}

type SqlcUpdate interface {
	GetUpdate() string
}

type SqlcDelete interface {
	GetDelete() string
}
