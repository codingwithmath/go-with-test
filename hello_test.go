package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Matheus")

	want := "Hello, Matheus!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}