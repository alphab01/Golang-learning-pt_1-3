package main

import (
	"maps"
	"testing"
)

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
	if !maps.Equal(mem, need) {
		t.Error("err")
	}
}

func TestStrparse3(t *testing.T) {
	mem := make(map[string]int)
	strparse2("привет мир", mem)
	need := map[string]int{"привет": 1, "мир": 1}
	if !maps.Equal(mem, need) {
		t.Error("err")
	}
}
