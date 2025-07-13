package data

import (
	"log/slog"
	"strings"
)

type StandardData struct {
	Name   string
	Logger *slog.Logger
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

func (sd *StandardData) GetLogger() *slog.Logger {
	return sd.Logger
}
