package data

import (
	"strings"
)

type StandardData struct {
	Name string
}

func (sd *StandardData) GetName() string {
	return sd.Name
}

func (sd *StandardData) GetAbbv() string {
	abbv := ""
	for _, c := range sd.GetName() {
		if rune('A') <= c && c <= rune('Z') {
			abbv += string(c)
		}
	}
	s := strings.ToLower(abbv)
	return s
}

func (sd *StandardData) GetLowerName() string {
	return strings.ToLower(sd.GetName())
}

func (sd *StandardData) GetStandardStruct() *StandardData {
	return sd
}
