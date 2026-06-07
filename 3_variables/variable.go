package main

import "fmt"

func main() {
	// var name string = "sajid"

	// When we don't explicitly specify the data type,
	// Go automatically infers (detects) the type from the assigned value.
	//it is for all the type not only for string
	// var name = "golang"

	// var isAdult bool = true
	// var isAdult = true

	// as we can see when we wrote int we are getting int , int32 , int64 , but we only write int , Go chooses the size of 'int' based on the underlying system architecture.
	// var age int = 45

	// shorthand syntax
	// name := "golang"

	// well the best practice is to use the type  because in some cases you assign first but use somewhere else

	// var name string

	// name = "majid"

	// var price float32 = 50.6

	price := 50.6

	fmt.Println(price)
}
