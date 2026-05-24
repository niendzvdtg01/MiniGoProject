package main

import "fmt"

func swap(a, b *int) (int, int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
	return *a, *b
}
func main() {
	a := 3
	b := 4
	swap(&a, &b)
	fmt.Println(a, b)
}
