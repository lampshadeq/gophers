package main

import "fmt"

func GetSum(nums ...int) {
  sum := 0
  for _, value := range nums {
    sum += value
  }
  fmt.Println("The sum is", sum)
}

func main() {
  fmt.Println("Getting sum for first 10 natural numbers.")
  GetSum(1,2,3,4,5,6,7,8,9,10)
  
  fmt.Println("Getting sum for integers 90 to 100")
  GetSum(90,91,92,93,94,95,96,97,98,99,100)
}