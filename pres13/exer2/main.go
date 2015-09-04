package main

import "fmt"

func printNums(nums ...int) {
  fmt.Println(nums)
}

func main() {
  firstSlice := []int{1, 3, 5, 7}
  printNums(firstSlice...)
  
  secondSlice := []int{5, 3, 35, 8, 6, 68, 14, 12, 1214}
  printNums(secondSlice...)
}