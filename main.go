package main

import (
  "fmt"
  "os"
  "unicode/utf8"
)

func strparse(d string, mem map[string]int) {
  s := ""
  for i := 0; i < utf8.RuneCountInString(d); i++ {
    if (string(d[i]) != " " && string(d[i]) != "\n") {
      s += string(d[i])
    } else if (string(d[i]) == " "){
      mem[s]++
      s = ""
    } else if (string(d[i]) == "\n") {
      mem[s]++
      s = ""
      mem["\n"]++
    }
  }
  mem[s]++
}

func howmuchsym(a string) int {
  s := 0
  for range a {
    s++
  }
  return s
}

func info(mem map[string]int) {
  s := 0
  for i, m := range mem {
    if (i != "" && i != " " && i != "\n") {
      s += len(i) * m
    }
  }
  fmt.Println("byte -", s)
  s = 0
  for i, m := range mem {
    if (i != "" && i != " " && i != "\n") {
      s += m
    }
  }
  fmt.Println("words -", s)
  s = 0
  for i, m := range mem {
    if (i != "" && i != " " && i != "\n") {
      s += (howmuchsym(i) * m)
    }
  }
  fmt.Println("syms -", s)
  s = 0
  for i, m := range mem {
    if (i == "\n") {
      fmt.Println("storcs -", m)
    }
  }
}

func main() {
  var s string
  fmt.Scan(&s)
  fmt.Println(s)
  f, _ := os.Open(s)
  d := make([]byte, 64)
  q, _ := f.Read(d)
  fmt.Println(string(d))
  mem := make(map[string]int)
  strparse(string(d[:q]), mem)
  //fmt.Println(mem["kirill"])
  info(mem)
}
