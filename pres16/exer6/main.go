package main

import "fmt"

func what(age int) (dogYears int) {
  dogYears = age * 7
  return
}

func main() {
  fmt.Println(what(15))
}