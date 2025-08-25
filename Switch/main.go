package main

import (
	"fmt"
	"time"
)

func main(){
	// simple switch

	// we do not have to write the break after the switch statement

	// i := 2

	// switch i {
	// case 1 :
	// 	fmt.Println("one")
	
	// case 2 :
	// 	fmt.Println("two")
	
	// case 3 :
	// 	fmt.Println("three")
	
	// default: 
	// 	fmt.Println("error")
	// }


	// **********************************
	// multiple condition switch

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's weekdays")

	}


}