package main
import ("fmt")

func main() {
  var a bool = true     
  var b int = 5         
  var c float32 = 3.14 
  var d string = "Hi!" 

  fmt.Println("Boolean: ", a)
  fmt.Println("Integer: ", b)
  fmt.Println("Float:   ", c)
  fmt.Println("String:  ", d)


   var x uint = 500
  var y uint = 4500
  fmt.Printf("Type: %T, value: %v", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)



//   example
  prices := [3]int{10,20,30}

  fmt.Println(prices[0])
  fmt.Println(prices[2])


  

}