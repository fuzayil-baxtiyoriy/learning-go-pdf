package main

import "fmt"

func main() {
	const x int64 = 10
	const (
		idKey = "id"
		nameKey = "name"
	)
	fmt.Println("Chapter 2")
	fmt.Println(idKey)
	fmt.Println(nameKey)
	fmt.Println(x)
}
