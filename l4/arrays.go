package main
import ("fmt")

func main() {
  var arr1 = [3]int{1,2,3}
  arr2 := [5]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)

//   method2 

  var arr3 = [...]int{1,2,3}
  arr4 := [...]int{4,5,6,7,8}

  fmt.Println(arr3)
  fmt.Println(arr4)
}

