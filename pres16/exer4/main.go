package main

import "fmt"

func getStr(name string, age string) string {
  return name + " is " + age + " years old."
}

func main() {
  name := "Herr"
  age := "25"
  
  fmt.Println(getStr(name, age))
}