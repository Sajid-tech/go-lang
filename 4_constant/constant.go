package main

import "fmt"

// const we can declare outside the function too

const age int = 30

var name int = 54

func main() {

	const name string = "golang"

	// name = "java"  cannot assign because it already assign

	fmt.Println(name)
	fmt.Println(age)

	// constant grouping

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
