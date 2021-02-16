package main

func Soma(numeros [5]int) int {
	soma := 0

	for _, numero := range numeros {
		soma += numero
	}

	return soma
}

func SomaSlice(numeros []int) int {
	soma := 0

	for _, numero := range numeros {
		soma += numero
	}

	return soma
}

func SomaTudo(numerosParaSomar ...[]int) []int {
	var somas []int

	for _, numeros := range numerosParaSomar {
		somas = append(somas, SomaSlice(numeros))
	}

	return somas
}

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
	var somas []int

	for _, numeros := range numerosParaSomar {
		if len(numeros) == 0 {
			somas = append(somas, 0)
		} else {
			final := numeros[1:]
			somas = append(somas, SomaSlice(final))
		}
	}

	return somas
}
