package main

import "fmt"

func fizzbuzz(limit int) {
  //I am ignoring 0 as it is always divisible XD
  i := 1
  for i <= limit{
    if i % 3 == 0 && i % 5 == 0{
      fmt.Println("fizzbuzz")
    }
    else if i % 5 == 0{
      fmt.Println("buzz")
    }
    else if i % 3 == 0{
      fmt.Println("fizz")
    }
    else{
      fmt.Println(i)
    }
    i++
  }
}

func main() {
  fizzbuzz(100)
  // display output in the next line
  //fmt.Print("Enter limit: ")
  
  // var then variable name then variable type
  //var limit int
  
  // Taking input from user
  //fmt.Scanln(&limit)
  //fizzbuzz(limit)
}
