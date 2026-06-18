package main

import "fmt"

func main() {
	age := 12

	// if age >= 18 {
	// 	fmt.Println("person is an adult")
	// } else {
	// 	fmt.Println("person is not an adult")
	// }

	if age >= 18 {
		fmt.Println("person is an adult")
	} else if age >= 13 && age < 18 {
		fmt.Println("person is a teenager")
	} else {
		fmt.Println("person is a child")
	}

	var role = "admin"
	var hasPermissions = false

	if role == "admin" && hasPermissions {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	// variable can be decarle directly inside if

	if age := 18; age >= 18 {
		fmt.Println("person is adult")
	} else {
		fmt.Println("not an adult")
	}

	// go does not have ternary operator, you will have to use normal if else

}
