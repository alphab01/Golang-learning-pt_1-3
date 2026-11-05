package main

import (
  "fmt"
  "os"
  "unicode/utf8"
  "strings"
  "strconv"
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
    } else if (string(d[i]) == " " || string(d[i]) == "," || string(d[i]) == "." || string(d[i]) == "!" || string(d[i]) == "?" || string(d[i]) == ";" || string(d[i]) == ":"){
      mem[strings.ToLower(s)]++
      s = ""
    } else if (string(d[i]) == "\n") {
      mem[strings.ToLower(s)]++
      s = ""
      mem["\n"]++
    }
  }
  mem[strings.ToLower(s)]++
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
      } else if (q[j].f == q[j + 1].f) {
        d1 := 0
        for k := 0; k < q[j].s && k < q[j + 1].s; k++ {
          if ((q[j].s)[k] < (q[j + 1].s)[k]) {
            d1 = 1
          } else if ((q[j].s)[k] > (q[j + 1].s)[k]) {
            d1 = 2
          }
        }
        if (d1 == 0) {
          if (utf8.RuneCountInString(q[j].s) < utf8.RuneCountInstring(q[j + 1].s)) {
            d1 = 1
          } else {
            d1 = 2
          }
        }
        if (d1 == 2) {
          fmt.Println("Ya ustal, boss")
        }
      }
    }
  }
  for i := 0; i < d && i < len(q); i++ {
    s := fmt.Sprintf("%d.", i + 1)
    fmt.Println(s, q[i].f, " - ", q[i].s)
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

func getop(read string) (op string, n int, list []string){
  s := ""
  fmt.Println(read)
  for i := 0; i < len(read); i++ {
    if (string(read[i]) == " ") {
      if (s == "topn") {
        op = s
        j := i + 1
        s = ""
        for j < len(read) {
          if (string(read[j]) != " ") {
            s += string(read[j])
          } else {
            n, _ = strconv.Atoi(s)
            k := j + 1
            s = ""
            for k < len(read) {
              if (string(read[k]) == " ") {
                list = append(list, s)
                s = ""
              } else {
                s += string(read[k])
              }
              k++
            }
            list = append(list, s)
            return
          }
          j++
        }
      } else {
        n = 0
        k := i + 1
        s = ""
        for k < len(read) {
          if (string(read[k]) == " ") {
            list = append(list, s)
            s = ""
          } else {
            s += string(read[k])
          }
          k++
        }
        list = append(list, s)
        return
      }
    } else {
      s += string(read[i])
    }
  }
  n = 0
  op = "notgiven"
  return
}

func main() {
  s := ""
  var read string
  for {
    _, err := fmt.Scan(&read)
    if (err != nil) {
      break
    }
    s += read
    s += " "
  }
  fmt.Println(s)
  op, n, list := getop(s)
  fmt.Println(op, n)
  if (op == "topn") {
    for i := range list {
      f, err := os.Open(list[i])
      if (err != nil) {
        fmt.Println("no file:(((")
      }
      d := make([]byte, 1000)
      q, _ := f.Read(d)
      mem := make(map[string]int)
      strparse(string(d[:q]), mem)
      fmt.Println(list[i], ":\n")
      topn(n, mem)
      fmt.Println()
      f.Close()
    }
  } else if (op == "stats"){
    for i := range list {
      f, err := os.Open(list[i])
      if (err != nil) {
        continue
      }
      d := make([]byte, 1000)
      q, _ := f.Read(d)
      mem := make(map[string]int)
      strparse(string(d[:q]), mem)
      fmt.Println(list[i], ":\n")
      info(mem)
      fmt.Println()
      f.Close()
    }
  }
  /*var s, reg string
  fmt.Scan(&s, &reg)
  //fmt.Println(s)
  f, err := os.Open(s)
  if (err != nil) {
    fmt.Println("no file:((")
  } else {
    d := make([]byte, 1000)
    q, _ := f.Read(d)
    //fmt.Println(string(d))
    mem := make(map[string]int)
    strparse(string(d[:q]), mem)
    //fmt.Println(mem["kirill"])
    if (reg == "stats") {
      info(mem)
    } else {
      var c int
      fmt.Scan(&c)
      topn(c, mem)
    }
  }*/
}
