package goarea

import "math"

//PI É UMA PROPORÇÃO NUMÉRICA
const Pi = 3.1415

// Circ é responsavel por receber a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é responsável por calcular a aárea do triangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Não visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
