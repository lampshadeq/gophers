package main

import (
  "fmt"
  "strconv"
)

func what(name string, age int) (int, bool) {
  return age * 7, age > 25
}

func main() {
  name := "Jaaaane"
  age := 26
  
  dog, old := what(name, age)
  
  if old {
    fmt.Println(name + " is " + strconv.Itoa(dog) + " in dog years and is old")
  } else {
    fmt.Println(name + " is " + strconv.Itoa(dog) + " in dog years and is not old")
  }
}