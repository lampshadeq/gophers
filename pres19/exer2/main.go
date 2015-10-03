package main

import "fmt"

type Animal struct {
  sound string
}

type Dog struct {
  Animal
  friendly bool
}

type Cat struct {
  Animal
  annoying bool
}

func specs(a interface{}) {
  fmt.Println(a)
}

func main() {
  fido := Dog {Animal {"weof"}, true}
  fifi := Cat {Animal {"meow"}, true}
  shadow := Dog { Animal {"moo"}, true}
  specs(fido)
  specs(fifi)
  specs(shadow)
}