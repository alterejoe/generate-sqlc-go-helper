package interfaces

import "github.com/dave/dst"

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
	GetReceiver() *dst.FieldList
	GetFunctionParams() *dst.FieldList
	GetResults() *dst.FieldList
	GetBody() *dst.BlockStmt
}
