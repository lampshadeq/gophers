package main

import (
  "encoding/csv"
  "fmt"
  "io"
  "log"
  "os"
  "strconv"
)

type state struct {
  id int
  name string
  abbr string
  region string
}

func parseState(columns map[string]int, record []string) (*state, error) {
  id, err := strconv.Atoi(record[columns["id"]])
  name := record[columns["name"]]
  abbr := record[columns["abbr"]]
  region := record[columns["region"]]
  
  if err != nil {
    return nil, err
  }
  
  return &state{
    id:id,
    name:name,
    abbr:abbr,
    region:region
  }, nil
}

func main() {
  f, err := os.Open("table.csv")
  if err != nil {
    log.Fatalln(err)
  }
  defer f.Close()
  
  csvReader := csv.NewReader(f)
  columns := make(map[string]int)
  stateLookup := map[string]*state{}
  
  for rowCount := 0; ; rowCount++ {
    record, err := csvReader.Read()
    
    if err == io.EOF {
      break
    } else if err != nil {
      log.Fatalln(err)
    }
    
    if rowCount == 0 {
      for idx, column := range record {
        columns[column] = idx
      }
    } else {
      state, err := parseState(columns, record)
      if err != nil {
        log.Fatalln(err)
      }
      
      stateLookup[state.abbr] = state
    }
  }
  
  if len(os.Args) < 2 {
    log.Fatalln("Expected state abbr")
  }
  
  abbr := os.Args[1]
  state, ok := stateLookup[abbr]
  if !ok {
    log.Fatalln("Invalid state abbr")
  }
  
  fmt.Println(`
  <html>
  <head></head>
  <body>
    <table>
      <tr>
        <th>Abbreviation</th>
        <th>Name</th>
      </tr>`)
      
  fmt.Println(`
      <tr>
        <td>` + state.abbr + `</td>
        <td>` + state.name + `</td>
      </tr>`)
      
  fmt.Println(`
    </table>
  </body>
 </html>`)
}