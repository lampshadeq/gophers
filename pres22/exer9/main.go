package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
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
    if len(line) > 0 {
      fmt.Println(">>>", strings.ToUpper(line[0:1] + line[1:], "\n")
    }
  }
}