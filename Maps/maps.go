package main

import "fmt"

// maps-> hash, object , dictionaries ,
// same as in c++ hashmap


func main() {

	// creating map

	//map_name:= make(map[key]value) 

	// m1:= make(map[string] int)
	// m2:= make(map[string] string)

	// m1["age"] = 20;
	// m1["umar"] = 23;
	// m2["name"] = "piyu";
	// m2["naam"] = "piyush";


	// fmt.Println(m1)
	// fmt.Println(m2)
	// fmt.Println(len(m1))
	// fmt.Println(len(m2))
	
	// // delete keyword in map
	// delete(m1,"age")
	// fmt.Println(m1)
	
	// // to clear a map
	// clear(m1)
	// fmt.Println(m1)


	// creating the map without using the make keyword

	// m:= map[string]int{"name":20 , "price":21}
	// fmt.Println(m)




	var m=  map[string]int{"name":20 , "price":21}
	fmt.Println(m)


}
