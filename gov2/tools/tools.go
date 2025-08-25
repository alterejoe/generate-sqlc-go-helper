package tools

import "strings"

func GetAbbv(name string) string {
	var capitalleters []string
	for _, char := range name {
		if char >= 'A' && char <= 'Z' {
			capitalleters = append(capitalleters, string(char))
		}
	}
	return strings.ToLower(strings.Join(capitalleters, ""))
}
