package area

import "math"

const Pi = 3.14

func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

func Ret(base, altura float64) float64 {
	return base * altura
}

func _Triangulo(base, altura float64) float64 {
	return (base * altura) / 2
}
