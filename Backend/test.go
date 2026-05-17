package main

import (
	"fmt"
	"strings"
)

func main() {
	var cc string = "abc,cde,asjdja"
	afterTest := strings.Split(cc, ",")
	joinTest := strings.Join(afterTest, "")
	fmt.Println(afterTest, joinTest)
}
