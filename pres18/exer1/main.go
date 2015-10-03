package main

import "fmt"

type person struct {
  fname string
  lname string
  age int
}

func (p *person) changeAge(newAge int) {
  p.age = newAge
}

func main() {
  p1 := person{"House", "Targaryen", 1000}
  fmt.Println(p1.age)
  p1.changeAge(0)
  fmt.Println(p1.age)
}