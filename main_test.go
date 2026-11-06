package main

import (
	"testing"
)

func mapcheck(a map[string]int, b map[string]int) bool {
	for i, m := range a {
		if b[i] != m && i != "" {
			return false
		}
	}
	for i, m := range b {
		if a[i] != m && i != "" {
			return false
		}
	}
	return true
}

func TestHowmuchsym(t *testing.T) {
	got := howmuchsym("kirill")
	if got != 6 {
		t.Error("err")
	}
}

func TestStrparse2(t *testing.T) {
	mem := make(map[string]int)
	strparse2("kirill kirill\n", mem)
	need := map[string]int{"kirill": 2, "\n": 1}
	if !mapcheck(mem, need) {
		t.Error("err")
	}
}

func TestStrparse3(t *testing.T) {
	mem := make(map[string]int)
	strparse2("привет мир", mem)
	need := map[string]int{"привет": 1, "мир": 1}
	if !mapcheck(mem, need) {
		t.Error("err")
	}
}
