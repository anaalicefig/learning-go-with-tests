package main

import "testing"

func TestHello(t *testing.T) {
	response := Hello("Ana")
	expected := "Hello Ana"

	if response != expected {
		t.Errorf("Response '%s', expected '%s'", response, expected)
	}
}
