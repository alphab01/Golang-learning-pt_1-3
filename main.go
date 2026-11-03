package main

import (
  "fmt"
  "os"
)

func strparse(d string, mem map[string]int) {
  s := ""
  for i := 0; i < len(d); i++ {
    if (string(d[i]) != " ") {
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
  mem := make(map[string]int)
  strparse(string(d), mem)
  fmt.Println(mem["kirill"])
}
