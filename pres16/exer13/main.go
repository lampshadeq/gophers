package main

import "fmt"

func marel(nums []int, callback func(int)) {
  for _, n := range nums {
    callback(n)
  }
}

func main() {
  marel([]int{1, 2, 3, 4}, func(n int) {
    fmt.Println(n)
  })
}