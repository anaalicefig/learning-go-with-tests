package inteiros

import (
	"fmt"
	"testing"
)

func TestAdicionador(t *testing.T) {
	soma := Adiciona(2, 2)
	esperado := 4

	if soma != esperado {
		t.Errorf("Esperado '%d', resultado '%d'", esperado, soma)
	}
}

func ExampleAdiciona() {
	soma := Adiciona(2, 4)
	fmt.Println(soma)
	// Output: 6
}