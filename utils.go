package main

import (
	"github.com/go-playground/validator/v10"
	"log/slog"
	"strings"
)

func ParseSlogLevel(s string) (slog.Level, error) {
	var level slog.Level
	var err = level.UnmarshalText([]byte(s))
	return level, err
}

var PassportNumber validator.Func = func(fl validator.FieldLevel) bool {
	passport := strings.Fields(fl.Field().String())
	if len(passport) != 2 {
		return false
	}
	return len(passport[0]) == 4 && len(passport[1]) == 6
}
