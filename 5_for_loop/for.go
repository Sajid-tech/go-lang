package main

import "fmt"

// for looping in go we have only for loop that it (we don't have while loop ) we do while loop use for loop
// only construct in go for looping
func main() {

	//while loop
	i := 1
	for i <= 5 { //no bracket nned in condition
		fmt.Println(i)
		i = i + 1
	}

	// infine loop
	// for {
	// 	//we use without fmt only for debuging and learning : builtin function
	// 	println("1")
	// }

	//classic for loop

	// for i := 0; i < 3; i++ {
	// 	// break  : break the iteration
	// 	if i == 2 {
	// 		continue // skip the iteration
	// 	}
	// 	fmt.Println(i)
	// }

	// range
	for i := range 3 {
		fmt.Println(i)
	}

}
