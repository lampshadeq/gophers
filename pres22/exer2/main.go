package main

import (
  "fmt"
  "strings"
)

func wordCount(str string) map[string]int {
  var words []string = strings.Fields(str)
  m := make(map[string]int)
  
  for _, v := range words {
    m[v]++
  }
  
  return m
}

func main() {
  str := `I am seventeen mushrooms about four scores ago when Daenarys who the pink pato??`
  fmt.Println(wordCount(str))
}