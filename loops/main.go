package main

import "fmt"

// only for loop
func main() {

	// while loop
	// var i int = 1
	i:=1
	for i<=3{
		fmt.Println(i)
		i+=1
	}

	// infiniite loop
	// for{
	// 	println(1)
	// }

	// normal for loop
	for i:=0;i<3;i++{
		// break
		fmt.Println(i)
		continue 
	}

	//range in for loop

	for i := range 4{
		fmt.Println(i)
	}


}
