package main

import "testing"

func TestHello(t *testing.T) {
	assetMessage := func(t *testing.T, result, expected string) {
		t.Helper()

		if result != expected {
			t.Errorf("Result %s, should %s", result, expected)
		}
	}

	t.Run("Say hello to people", func(t *testing.T) {
		result := Hello("Ana", "spanish")
		expected := "Hello Ana"

		assetMessage(t, result, expected)
	})

	t.Run("Say 'Hello world' when an empty string is passed", func(t *testing.T) {
		result := Hello("", "spanish")
		expected := "Hello World"

		assetMessage(t, result, expected)
	})

	t.Run("Say hello in spanish", func(t *testing.T) {
		result := Hello("Diego", "spanish")
		expected := "Hola, Diego"

		assetMessage(t, result, expected)
	})
}
