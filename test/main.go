package main

import "fmt"

func main() {
	a := []byte{'b', 'b'}
	c := a[:0]
	fmt.Println(c, c == nil, len(c), cap(c))
	fmt.Println(a, c)
}
