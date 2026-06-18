package main

import (
	"fmt"
	"regexp"
)

func main() {
	testString := "ABCD.a"
	fmt.Println(camelToSnake(testString))
}

var (
	matchFirstCap = regexp.MustCompile("(.)[A-Z][a-z]+")
)

func camelToSnake(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	return snake
}
