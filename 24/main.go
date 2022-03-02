package main

import (
	"fmt"
	"math"
)

/*Задание:
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point
с инкапсулированными параметрами x,y и конструктором.
*/

type Point struct {
	x, y float64
}

func main() {
	// создаем две точки
	point1 := NewPoint(0, 0)
	point2 := NewPoint(1, -5)

	// вычисляем расстояние между ними и выводим его в stdout
	fmt.Printf("%.2f", PointDist(point1, point2))
}

// NewPoint конструктор структуры Point
func NewPoint(x, y float64) *Point {
	// задаем значения полей и возвращаем указатель на созданную структуру
	return &Point{x: x, y: y}
}

// PointDist вычисляет расстояние между point1 и point2
func PointDist(p1, p2 *Point) float64 {
	// вычисляем расстояние по формуле
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}
