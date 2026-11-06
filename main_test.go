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
