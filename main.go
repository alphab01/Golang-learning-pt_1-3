package main

import (
  "fmt"
  "os"
  "unicode/utf8"
)

type pairintstr struct {
  f int
  s string
}

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

func topn(d int, mem map[string]int) {
  var q []pairintstr
  for i, m := range mem {
    if (i != "" && i != " " && i != "\n") {
      p := pairintstr{f: m, s: i}
      q = append(q, p)
    }
  }
  for i := 0; i < len(q); i++ {
    for j := 0; j < len(q) - 1; j++ {
      if (q[j].f < q[j + 1].f) {
        p := pairintstr{f: q[j].f, s: q[j].s}
        q[j].f = q[j + 1].f
        q[j].s = q[j + 1].s
        q[j + 1].f = p.f
        q[j + 1].s = p.s
      }
    }
  }
  for i := 0; i < d && i < len(q); i++ {
    fmt.Println(q[i].f, " - ", q[i].s)
  }
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
  fmt.Println("strok -", mem["\n"])
  ma := 0
  str := ""
  for i, _ := range mem {
    if (i != "" && i != " " && i != "\n") {
      if (ma < utf8.RuneCountInString(i)) {
        ma = utf8.RuneCountInString(i)
        str = i
      }
    }
  }
  fmt.Println("MOST BIG -", str)
}

func main() {
  var s, reg string
  fmt.Scan(&s, &reg)
  //fmt.Println(s)
  if (reg == "stats") {
    f, _ := os.Open(s)
    d := make([]byte, 64)
    q, _ := f.Read(d)
    //fmt.Println(string(d))
    mem := make(map[string]int)
    strparse(string(d[:q]), mem)
    //fmt.Println(mem["kirill"])
    info(mem)
    topn(2, mem)
  }
}
