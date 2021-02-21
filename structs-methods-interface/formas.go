package main

import "math"

// Se o tipo que você passar combinar com o que a interface está esperando, o código será compilado.
type Forma interface {
	Area() float64
}

type Triangulo struct {
	Base float64
	Altura float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) * 0.5
}

type Retangulo struct {
	Largura float64
	Altura  float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}

func Area(retangulo Retangulo) float64 {
	return retangulo.Largura * retangulo.Altura
}
