package conversions

import (
	"fmt"
	"strings"

	"github.com/dave/dst"
)

type FromFuncDecl struct {
	Name    string
	Input   *dst.FuncDecl
	Params  *dst.FieldList
	Results *dst.FieldList
}

func (fgd *FromFuncDecl) GetName() string {
	fgd.Name = fgd.Input.Name.Name
	return fgd.Name
}

func (qmp *FromFuncDecl) GetLowerName() string {
	return strings.ToLower(qmp.Name)
}

func (qmp *FromFuncDecl) GetReturns() *dst.FieldList {
	return &dst.FieldList{
		List: []*dst.Field{
			{Type: dst.NewIdent("any")},
			{Type: dst.NewIdent("error")},
		},
	}
}

func (qmp *FromFuncDecl) GetQueryArgs(params *[]dst.Expr) {
	qmp.Params = qmp.Input.Type.Params
	if len(qmp.Params.List) > 1 {
		*params = append(*params, dst.NewIdent(fmt.Sprintf("*%s.%s", qmp.GetAbbv(), qmp.GetSecondArg())))
	}
}

func (qmp *FromFuncDecl) GetSecondArg() string {
	if len(qmp.Params.List) > 1 {
		t := qmp.Params.List[1].Type
		if !strings.Contains(fmt.Sprint(t), "Params") {
			return fmt.Sprint(qmp.Params.List[1].Names[0])
		} else {
			return "Params"
		}
	}
	return ""
}

func (qmp *FromFuncDecl) GetAbbv() string {
	abbv := ""
	for _, c := range qmp.Name {
		if rune('A') <= c && c <= rune('Z') {
			abbv += string(c)
		}
	}
	s := strings.ToLower(abbv)
	return s
}
func (qmp *FromFuncDecl) GetQueryAddErr(results *[]dst.Expr, function bool) {
	qmp.Results = qmp.Input.Type.Results
	if len(qmp.Results.List) > 1 {
		*results = append(*results, dst.NewIdent("err"))
	} else if !function {
		*results = append(*results, dst.NewIdent("nil"))
	}
}

func (qmp *FromFuncDecl) StructParams() *dst.Field {
	qmp.Params = qmp.Input.Type.Params
	if len(qmp.Params.List) > 1 {
		t := qmp.Params.List[1].Type
		if !strings.Contains(fmt.Sprint(t), "Params") {
			return &dst.Field{
				Names: qmp.Params.List[1].Names,
				Type:  dst.NewIdent(fmt.Sprint("*", t)),
			}
		} else {
			field := &dst.Field{
				Names: []*dst.Ident{dst.NewIdent("Params")},
				Type:  dst.NewIdent(fmt.Sprint("*db.", t)),
			}
			return field
		}
	}
	return nil
}
