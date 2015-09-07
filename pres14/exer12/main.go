package main

import (
  "fmt"
  "reflect"
)

func main() {
  myVar := "Yo?"
  myNum := 125.33
  myRune := ','
  
  fmt.Println("myVar\t", reflect.TypeOf(myVar), "\t", myVar)
  fmt.Println("myNum\t", reflect.TypeOf(myNum), "\t", myNum)
  fmt.Println("myRune\t", reflect.TypeOf(myRune), "\t\t", myRune)
}