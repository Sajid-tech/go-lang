package main

import (
	"fmt"
	"time"
)

func main() {

	//simple switch  , if we we have multiple conditon ther we need to use  switch construct

	// no need to add break after case and alsd default is also optional

	i := 9

	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("other")
	}

	// multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its a weekend")
	default:
		fmt.Println("its a workday")
	}

	// type switch

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println(("its an integer"))
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("its a  boolean")
		default:
			fmt.Println("other", t)
		}
	}

	whoAmI("sajid")
	whoAmI(4)
	whoAmI(3.5)
	whoAmI(false)

}
