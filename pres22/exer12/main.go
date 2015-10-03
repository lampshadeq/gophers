package main

import (
  "bufio"
  "fmt"
  "io"
  "log"
  "os"
  "strings"
)

func wordCount(rdr io.Reader) map[string]int {
  counts := map[string]int{ }
  scanner := bufio.NewScanner(rdr)
  scanner.Split(bufio.ScanWords)
  
  for scanner.Scan() {
    word := scanner.Text()
    word = strings.ToLower(word)
    word = strings.Replace(word, ",", "", 1)
    word = strings.Replace(word, ".", "", 1)
    
    counts[word]++
  }
  
  return counts
}

func main() {
  src, err := os.Open("sample.txt")
  if err != nil {
    log.Fatalln(err)
  }
  defer src.Close()
  
  counts := wordCount(src)
  fmt.Println(counts["Tyrell"])
}