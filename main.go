package main

import (
  "fmt"
  "os"
  "unicode/utf8"
)

func removeslashn(d string) string {
  s := ""
  for i := 0; i < utf8.RuneCountInString(d); i++ {
    if (string(d[i]) != "\n") {
      s += string(d[i])
    } else {
      s += " "
    }
  }
  return s
}

func strparse(d string, mem map[string]int) {
  s := ""
  d = removeslashn(d)
  for i := 0; i < utf8.RuneCountInString(d); i++ {
    if (string(d[i]) != " ") {
      s += string(d[i])
    } else {
      if (s != "" && s != " ") {
        mem[s]++
      }
      s = ""
    }
  }
  if (s != " " && s != "") {
    mem[s]++
  }
}

func info(mem map[string]int) {
  s := 0
  for i, m := range mem {
    if (!(i == "" || i == " ")) {
      s += len(i) * m
    }
  }
  fmt.Println("byte - ", s)
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
  //fmt.Println(mem["kirill"])
  info(mem)
}
