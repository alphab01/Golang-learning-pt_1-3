package main

import (
  "fmt"
  "os"
)

func main() {
  var s string
  fmt.Scan(&s)
  fmt.Println(s)
  f, _ := os.Open(s)
  d := make([]byte, 64)
  f.Read(d)
  fmt.Println(string(d))
}
