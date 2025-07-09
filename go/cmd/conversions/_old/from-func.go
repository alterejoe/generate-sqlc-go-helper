package conversions

import (
	"fmt"
	"strings"

	"github.com/dave/dst"
)

type FromFuncDeclQuery struct {
	Input   *dst.FuncDecl
	Results *dst.FieldList
	Name    string
	QueryStruct
}

func (fgd *FromFuncDeclQuery) Initialize() {
	fgd.Name = fgd.Input.Name.Name
	fgd.Results = fgd.Input.Type.Results
	fgd.QueryStruct.Params = fgd.Input.Type.Params
	fgd.QueryStruct.Initialize(fgd)
}
func (fgd *FromFuncDeclQuery) GetName() string {
	return fgd.Name
}

func (qmp *FromFuncDeclQuery) GetLowerName() string {
	return strings.ToLower(qmp.Name)
}

func (qmp *FromFuncDeclQuery) GetReturns() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{Type: dst.NewIdent("any")},
			{Type: dst.NewIdent("error")},
		},
	}
}

func (qmp *FromFuncDeclQuery) GetQueryArgs(params *[]dst.Expr) {
	qmp.QueryStruct.Params = qmp.Input.Type.Params
	if len(qmp.QueryStruct.Params.List) > 1 {
		*params = append(*params, dst.NewIdent(fmt.Sprintf("*%s.%s", qmp.GetAbbv(), qmp.GetSecondArg())))
	}
}

func (qmp *FromFuncDeclQuery) GetSecondArg() string {
	if len(qmp.QueryStruct.Params.List) > 1 {
		t := qmp.QueryStruct.Params.List[1].Type
		if !strings.Contains(fmt.Sprint(t), "Params") {
			return fmt.Sprint(qmp.QueryStruct.Params.List[1].Names[0])
		} else {
			return "Params"
		}
	}
	return ""
}

func (qmp *FromFuncDeclQuery) GetAbbv() string {
	abbv := ""
	for _, c := range qmp.Name {
		if rune('A') <= c && c <= rune('Z') {
			abbv += string(c)
		}
	}
	s := strings.ToLower(abbv)
	return s
}
func (qmp *FromFuncDeclQuery) GetQueryAddErr(results *[]dst.Expr, function bool) {
	if len(qmp.Results.List) > 1 {
		*results = append(*results, dst.NewIdent("err"))
	} else if !function {
		*results = append(*results, dst.NewIdent("nil"))
	}
}
