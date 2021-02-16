package interacao

import (
	"strings"
	"testing"
)

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a", 10)
	esperado := strings.Repeat("a", 10)

	if repeticoes != esperado {
		t.Errorf("esperado '%s' mas objete '%s'", esperado, repeticoes)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 10)
	}
}
