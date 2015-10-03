package main

import (
  "fmt"
  "os"
  "bufio"
  "log"
)

func main() {
  src, err := os.Open("init.txt")
  if err != nil {
    log.Printf("Could not open file: %v", err)
  }
  defer src.Close()
  
  scanner := bufio.NewScanner(src)
  for scanner.Scan() {
    line := scanner.Text()
    fmt.Println(">>>", line)
  }
}