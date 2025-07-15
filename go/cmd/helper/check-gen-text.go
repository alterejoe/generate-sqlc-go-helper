package helper

import (
	"fmt"

	"github.com/dave/dst"
)

func CheckGenText(t dst.Expr) bool {
	// switch v := fmt.Sprintf(t.Type), v {
	switch v := fmt.Sprintf("%s", t); v {
	case "&{pgtype Text {{None [] [] None} []}}", "string", "&{<nil> bool {{None [] [] None} [] []}}", "&{<nil> string {{None [] [] None} [] []}}":
		return false
	default:
		return true
	}
}

func ToStandardReturnType(t *dst.Expr) string {
	switch fmt.Sprint(*t) {
	case "string", "&{pgtype Text {{None [] [] None} []}}":
		return "string"
	case "int64", "int32", "&{pgtype Int4 {{None [] [] None} []}}":
		return "int"
	case "bool", "&{pgtype Bool {{None [] [] None} []}}":
		return "bool"
	case "&{pgtype Float8 {{None [] [] None} []}}":
		return "float64"
	case "&{pgtype Timestamp {{None [] [] None} []}}", "&{pgtype Timestamptz {{None [] [] None} []}}":
		return "time.Time"
	case "&{<nil> byte {{None [] [] None} [] []}}":
		return "[]byte"
	case "&{<nil> string {{None [] [] None} [] []}}":
		return "[]string"
	case "&{<nil> float64 {{None [] [] None} [] []}}":
		return "[]float"
	case "&{<nil> int32 {{None [] [] None} [] []}}":
		return "[]int"
	case "&{<nil> bool {{None [] [] None} [] []}}":
		return "[]bool"
	case "&{pgtype UUID {{None [] [] None} []}}":
		return "uuid.UUID"
	default:
		fmt.Println("Unknown return type", *t)
		return "NOTHING"
	}
}

func ToStandardReturn(t *dst.Expr) dst.Expr {
	switch fmt.Sprint(*t) {
	case "string", "int64", "int32", "bool":
		return *t
	case "&{pgtype Bool {{None [] [] None} []}}":
		return dst.NewIdent("false")
	case "&{pgtype Int4 {{None [] [] None} []}}", "&{pgtype Float8 {{None [] [] None} []}}":
		return dst.NewIdent("-1")
	case "&{pgtype Timestamp {{None [] [] None} []}}", "&{pgtype Timestamptz {{None [] [] None} []}}":
		return dst.NewIdent("time.Time{}")
	case "&{<nil> byte {{None [] [] None} [] []}}":
		return dst.NewIdent("[]byte{}")
	case "&{<nil> string {{None [] [] None} [] []}}":
		return dst.NewIdent("[]string")
	case "&{<nil> float64 {{None [] [] None} [] []}}":
		return dst.NewIdent("[]float{}")
	case "&{<nil> int32 {{None [] [] None} [] []}}":
		return dst.NewIdent("[]int{}")
	case "&{<nil> bool {{None [] [] None} [] []}}":
		return dst.NewIdent("[]bool{}")
	case "&{pgtype UUID {{None [] [] None} []}}":
		return dst.NewIdent("uuid.UUID{}")
	case "&{pgtype Text {{None [] [] None} []}}":
		return dst.NewIdent("\"\"")
	default:
		fmt.Println("Unknown return", *t)
		return dst.NewIdent("\"\"")
	}
}
