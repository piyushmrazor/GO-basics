package main

import (
	"fmt"
	"net/http"
)

// Declaring the variable globally
//var array = []string{"Watch here the youtuber channel", "Watch go crash course", "Learning the Golang"}

// Defining the parameters in the functions with their datatypes
func PrintTasks(array []string) {
	for idx, i := range array {
		fmt.Println(idx+1, ".", i)
	}
}

// Here the list for the addded new task
func addTasks(array *[]string, newTask string) {
	*array = append(*array, newTask)
	PrintTasks(*array)
}

func main() {
	//var name = "######## Welcome #########"
	fmt.Println("Watch go crash course")
	//fmt.Println("Watch here the youtuber channel")
	//fmt.Println("Learning the Golang")
	//
	//fmt.Println("")

	//Slice arrays
	//var array = []string{"Watch here the youtuber channel", "Watch go crash course", "Learning the Golang"}
	////fmt.Println(array)
	//fmt.Println("Function is calling ")
	////Calling the function
	//
	////we can pass as many parametes as possible in the functions
	////PrintTasks(array)
	//var newtasks = "hey there"
	//addTasks(&array, newtasks)
	//
	//fmt.Println("Function is executed ")
	//PrintTasks(array)
	//for idx, i := range array {
	//	fmt.Println(idx+1, ".", i)
	//}

	//Making the http server
	//*************************************
	http.ListenAndServe(":8080", nil)

}
