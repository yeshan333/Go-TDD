package main

import "testing"

func TestNewFunc(t *testing.T) {
	got := NewFunc()

	want := "Golang is Good!"

	if *got != want {
		t.Errorf("want %s, got %s", want, *got)
	}
}
