package main

import (
	"fmt"
	"math"
)

// Структура для представления точки.
type Point struct {
	x float64
	y float64
}

// Конструктор для создания новой точки.
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Метод для вычисления расстояния между двумя точками.
func (p *Point) Distance(anotherPoint *Point) float64 {
	return math.Sqrt(math.Pow(anotherPoint.x-p.x, 2) + math.Pow(anotherPoint.y-p.y, 2))
}

func main() {
	p1 := NewPoint(2, 3)
	p2 := NewPoint(4, 6)

	distance := p1.Distance(p2)
	fmt.Printf("Расстояние между точками: %f\n", distance)
}
