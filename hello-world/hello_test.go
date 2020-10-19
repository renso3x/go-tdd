package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Romeo")
	want := "Hello, Romeo"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
