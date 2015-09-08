package main

import "fmt"

type customer struct {
  name string
  balance float64
}

func main() {
  myCustomer := customer{"Mace", 15263.35}
  
  fmt.Println(myCustomer.name)
  fmt.Println(myCustomer.balance)
  
  myCustomer.balance = 10.52
  
  fmt.Println(myCustomer.balance)
}