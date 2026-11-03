package main

import (
  "fmt"
  "os"
)

func strparse(d string, mem [string]int) {
  s := ""
  for i := 0; i < len(d); i++ {
    if (d[i] != " ") {
      s += string(d[i])
    } else {
      mem[s]++
      s = ""
    }
  }
  mem[s]++
}

func main() {
  var s string
  fmt.Scan(&s)
  fmt.Println(s)
  f, _ := os.Open(s)
  d := make([]byte, 64)
  f.Read(d)
  fmt.Println(string(d))
  mem := make([string]int)
}
