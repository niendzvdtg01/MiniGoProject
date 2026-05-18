package utils

import (
	"regexp"
	"strings"
)

var (
	matchFirstCap = regexp.MustCompile("(.)[A-Z][a-z]+")
	matchAllcap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func camelToSnake(str string) string {
	snake := matchAllcap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllcap.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(snake)
}
