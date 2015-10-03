package main

import (
  "bufio"
  "fmt"
  "io"
  "log"
  "os"
)

func longestWord(rdr io.Reader) string {
  longWord := ""
  scanner := bufio.NewScanner(rdr)
  scanner.Split(bufio.ScanWords)
  
  for scanner.Scan() {
    word := scanner.Text()
    if len(word) > len(longWord) {
      longWord = word
    }
  }
  
  return longWord
}

func main() {
  src, err := os.Open("sample.txt")
  if err != nil {
    log.Fatalln(err)
  }
  defer src.Close()
  
  fmt.Println(longestWord(src))
}