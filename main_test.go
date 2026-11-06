package main

import (
  "testing"
)

func TestHowmuchsym(t *testing.T) {
  got := howmuchsym("kirill")
  if (got != 6) {
    t.Error("err")
  }
}

func TestStrparse2(t *testing.T) {
  mem := make(map[string]int)
  strparse2("kirill kirill\n", mem)
  need := map[string]int{"kirill": 2, "\n": 1}
  if (mem != need) {
    t.Error("err")
  }
}
