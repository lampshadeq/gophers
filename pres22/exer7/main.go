package main

import (
  "os"
  "log"
  "io/ioutil"
)

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    log.Fatalln("Unable to open file", err.Error())
  }
  defer f.Close()
  
  nf, err := os.Create("newFile.txt")
  if err != nil {
    log.Fatalln("Unable to create file", err.Error())
  }
  
  bs, err := ioutil.ReadAll(f)
  if err != nil {
    log.Fatalln("Unable to read file", err.Error())
  }
  
  _, err = nf.Write(bs)
  if err != nil {
    log.Fatalln("Broken?", err.Error())
  }
}