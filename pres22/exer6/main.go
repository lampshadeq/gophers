package main

import (
  "os"
  "log"
  "io/ioutil"
  "fmt"
)

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    log.Fatalln("Unable to open file", err.Error())
  }
  defer f.Close()
  
  bs, err := ioutil.ReadAll(f)
  if err != nil {
    log.Fatalln("Unable to read file")
  }
  
  str := string(bs)
  fmt.Println(str)
}